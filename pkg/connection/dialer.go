package connection

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/josh-silvas/dmux/internal/nlog"
	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
	"golang.org/x/term"
)

var l = nlog.NewWithGroup("connection")

type (
	// SSH structure to store contents about ssh connection.
	SSH struct {
		Client                *ssh.Client
		ProxyDialer           proxy.Dialer
		ConnectTimeout        int
		SendKeepAliveMax      int
		SendKeepAliveInterval int
		logging               bool
		logFile               string
	}
)

// NewDialer : Func to initialize the SSH struct and set up the ssh session.
func NewDialer(host, user, pass string) (*SSH, error) {
	c := new(SSH)

	timeout := 20
	if c.ConnectTimeout > 0 {
		timeout = c.ConnectTimeout
	}

	l.Debugf("Creating SSH connection to %s with user %s", host, user)

	// Parse SSH config file
	cfg, err := ssh_config.Decode(getSSHConfig())
	if err != nil {
		l.Warnf("Failed to parse SSH config: %v", err)
	}

	configUser := getConfigUser(cfg, host, user)
	if configUser != user {
		l.Debugf("Using user %s from SSH config instead of provided user %s", configUser, user)
	}

	// Create new ssh.ClientConfig{}
	config := &ssh.ClientConfig{
		User:    configUser,
		Auth:    []ssh.AuthMethod{ssh.Password(pass)},
		Timeout: time.Duration(timeout) * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			l.Debugf("Accepting host key for %s (%s)", hostname, remote)
			return nil
		},
		HostKeyAlgorithms: getHostKeyAlgorithmsFromConfig(cfg, host),
		Config: ssh.Config{
			KeyExchanges: getKeyExchangesFromConfig(cfg, host),
			Ciphers:      getCiphersFromConfig(cfg, host),
		},
	}

	// Get port from config or use default
	port := ssh_config.Get(host, "Port")
	if port == "" {
		port = "22"
	}
	hostPort := net.JoinHostPort(host, port)
	l.Debugf("Connecting to %s", hostPort)

	// Check for ProxyCommand in SSH config
	// ProxyCommand allows SSH connections to be established through an intermediate host
	// using a command specified in the SSH config. For example:
	//   Host target
	//     ProxyCommand ssh jumphost nc %h %p
	// This will connect to 'jumphost' first, then use netcat (nc) to establish
	// a connection to the final target host (%h) and port (%p).
	if proxyCmd := ssh_config.Get(host, "ProxyCommand"); proxyCmd != "" {
		l.Debugf("Found ProxyCommand in config: %s", proxyCmd)

		// Parse the target host and port. If the host contains a port (e.g., host:22),
		// we need to split them to properly substitute into the ProxyCommand.
		targetHost := host
		targetPort := port
		if strings.Contains(host, ":") {
			parts := strings.Split(host, ":")
			targetHost = parts[0]
			targetPort = parts[1]
		}

		// Replace the placeholders in the ProxyCommand:
		// %h - target hostname
		// %p - target port
		// This allows the SSH config to be dynamic based on the target.
		proxyCmd = strings.ReplaceAll(proxyCmd, "%h", targetHost)
		proxyCmd = strings.ReplaceAll(proxyCmd, "%p", targetPort)
		l.Debugf("Using expanded ProxyCommand: %s", proxyCmd)

		// Create a proxy dialer that will execute the ProxyCommand
		// The command's stdout/stdin will be used as the network connection
		c.ProxyDialer = &pxyCmd{command: proxyCmd}
	} else {
		// If no ProxyCommand is specified, use a direct connection
		l.Debug("No proxy configured, using direct connection")
		c.ProxyDialer = proxy.Direct
	}

	// Dial to host:port
	l.Debug("Establishing TCP connection...")
	netConn, err := c.ProxyDialer.Dial("tcp", hostPort)
	if err != nil {
		var proxyErr *net.OpError
		if errors.As(err, &proxyErr) {
			l.Debugf("Network operation error: %v", proxyErr.Err)
		}
		return nil, fmt.Errorf("failed to establish TCP connection: %w", err)
	}
	l.Debug("TCP connection established successfully")

	// Set TCP keepalive
	if tcpConn, ok := netConn.(*net.TCPConn); ok {
		if err := tcpConn.SetKeepAlive(true); err != nil {
			l.Warnf("Failed to set TCP keepalive: %v", err)
		}
		if err := tcpConn.SetKeepAlivePeriod(30 * time.Second); err != nil {
			l.Warnf("Failed to set TCP keepalive period: %v", err)
		}
	}

	// Create new ssh connection
	l.Debug("Starting SSH handshake...")
	sshConn, chans, reqs, err := ssh.NewClientConn(netConn, hostPort, config)
	if err != nil {
		err := netConn.Close()
		if err != nil {
			return nil, err
		}
		l.Debugf("SSH handshake failed: %v", err)
		return nil, fmt.Errorf("SSH handshake failed: %w", err)
	}
	l.Debug("SSH handshake completed successfully")

	// Create *ssh.Client
	c.Client = ssh.NewClient(sshConn, chans, reqs)
	l.Debug("SSH client created successfully")

	// Verify connection is working
	if _, _, err := c.Client.SendRequest("keepalive@openssh.com", true, nil); err != nil {
		err := c.Client.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("connection verification failed: %w", err)
	}

	return c, nil
}

// getSSHConfig reads the SSH config file
func getSSHConfig() io.Reader {
	home, err := os.UserHomeDir()
	if err != nil {
		l.Warnf("Failed to get home directory: %v", err)
		return strings.NewReader("")
	}

	configFile := filepath.Join(home, ".ssh", "config")
	data, err := os.ReadFile(configFile) // #nosec G304
	if err != nil {
		l.Warnf("Failed to read SSH config file: %v", err)
		return strings.NewReader("")
	}

	return bytes.NewReader(data)
}

// getConfigUser gets the user from SSH config with fallback
func getConfigUser(cfg *ssh_config.Config, host, defaultUser string) string {
	if cfg == nil {
		return defaultUser
	}
	user := ssh_config.Get(host, "User")
	if user == "" {
		return defaultUser
	}
	return user
}

// getHostKeyAlgorithmsFromConfig returns the configured host key algorithms
func getHostKeyAlgorithmsFromConfig(cfg *ssh_config.Config, host string) []string {
	if cfg == nil {
		return getDefaultHostKeyAlgorithms()
	}

	algos := ssh_config.Get(host, "HostKeyAlgorithms")
	if algos == "" {
		return getDefaultHostKeyAlgorithms()
	}
	return strings.Split(algos, ",")
}

func getDefaultHostKeyAlgorithms() []string {
	return []string{
		ssh.KeyAlgoRSA,
		ssh.KeyAlgoDSA,
		ssh.KeyAlgoECDSA256,
		ssh.KeyAlgoECDSA384,
		ssh.KeyAlgoECDSA521,
		ssh.KeyAlgoED25519,
	}
}

// getKeyExchangesFromConfig returns the configured key exchange algorithms
func getKeyExchangesFromConfig(cfg *ssh_config.Config, host string) []string {
	if cfg == nil {
		return getDefaultKeyExchanges()
	}

	kex := ssh_config.Get(host, "KexAlgorithms")
	if kex == "" {
		return getDefaultKeyExchanges()
	}
	return strings.Split(kex, ",")
}

func getDefaultKeyExchanges() []string {
	return []string{
		"diffie-hellman-group1-sha1",
		"diffie-hellman-group14-sha1",
		"diffie-hellman-group14-sha256",
		"ecdh-sha2-nistp256",
		"ecdh-sha2-nistp384",
		"ecdh-sha2-nistp521",
		"curve25519-sha256@libssh.org",
		"curve25519-sha256",
	}
}

// getCiphersFromConfig returns the configured ciphers
func getCiphersFromConfig(cfg *ssh_config.Config, host string) []string {
	if cfg == nil {
		return getDefaultCiphers()
	}

	ciphers := ssh_config.Get(host, "Ciphers")
	if ciphers == "" {
		return getDefaultCiphers()
	}
	return strings.Split(ciphers, ",")
}

func getDefaultCiphers() []string {
	return []string{
		"aes128-gcm@openssh.com",
		"chacha20-poly1305@openssh.com",
		"aes128-ctr",
		"aes192-ctr",
		"aes256-ctr",
		"aes128-cbc",
		"aes192-cbc",
		"aes256-cbc",
	}
}

// SendKeepAlive : Method on the SSH struct to send keepalive across the
// channel a given interval
func (c *SSH) SendKeepAlive(session *ssh.Session) {
	// Set the default to 30sec. If a struct keepalive interval
	// is defined, override the default.
	interval := 30
	if c.SendKeepAliveInterval > 0 {
		interval = c.SendKeepAliveInterval
	}

	// Set the default to 5sec. If a struct keepalive max interval
	// is defined, override the default.
	max := 5
	if c.SendKeepAliveMax > 0 {
		max = c.SendKeepAliveMax
	}

	i := 0
	for {
		if _, err := session.SendRequest("keepalive", true, nil); err == nil {
			i = 0
		} else {
			i++
		}

		if max <= i {
			if err := session.Close(); err != nil {
				l.Errorf("[session.Close::%s]", err)
			}
			return
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// RequestTTY requests the association of a pty with the session on the remote
// host. Terminal size is obtained from the currently connected terminal
func RequestTTY(session *ssh.Session) (err error) {
	m := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	// Get terminal window size
	fd := int(os.Stdout.Fd())
	w, h, err := term.GetSize(fd)
	if err != nil {
		return
	}

	// Attempt to get the terminal from environment
	// variables. Use `xterm` by default.
	t := os.Getenv("TERM")
	if len(t) == 0 {
		t = "xterm"
	}

	if err = session.RequestPty(t, h, w, m); err != nil {
		if e := session.Close(); e != nil {
			closeErr := e.Error()
			err = fmt.Errorf("[session.Close::%w] [%s]", e, closeErr)
		}
		return
	}

	// Terminal resize goroutine.
	winch := syscall.Signal(0x1c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, winch)
	go func() {
		for {
			s := <-ch
			switch s {
			case winch:
				fd := int(os.Stdout.Fd())

				w, h, err = term.GetSize(fd)
				if err != nil {
					l.Errorf("[term.GetSize::%s]", err)
				}

				if err := session.WindowChange(h, w); err != nil {
					l.Errorf("[session.WindowChange::%s]", err)
				}
			}
		}
	}()

	return
}

// SetLog : Set a logging path for this terminal session.
func (c *SSH) SetLog(hostName string) {
	prof, err := user.Current()
	if err != nil {
		l.Errorf("[user.Current::%s]", err)
		return
	}

	fullPath := fmt.Sprintf("%s/.log/dmux/ssh", prof.HomeDir)
	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			l.Errorf("[SetLog.Mkdir::%s]", err)
			return
		}
	}
	fileName := fmt.Sprintf("%s_%s", hostName, time.Now().Format("2006-01-02"))
	c.logging = true
	c.logFile = fmt.Sprintf("%s/%s.log", fullPath, fileName)
}

// Shell : SSH method to launch an interactive SSH shell.
func (c *SSH) Shell(session *ssh.Session) error {
	fd := int(os.Stdin.Fd())

	state, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}

	defer func() {
		if e := term.Restore(fd, state); e != nil {
			err = e
		}
	}()

	if err := c.setupShell(session); err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}

	go c.SendKeepAlive(session)

	return session.Wait()
}

func (c *SSH) setupShell(session *ssh.Session) error {
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if c.logging {
		if err := c.logger(session); err != nil {
			l.Errorf("[Error setting up log path for session::%s]", err)
		}
	}

	if err := RequestTTY(session); err != nil {
		return err
	}

	return nil
}

// logger : SSH method to open a logfile and log output from terminal session.
func (c *SSH) logger(s *ssh.Session) error {
	logfile, err := os.OpenFile(c.logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("[os.OpenFile::%w]", err)
	}

	buf := new(bytes.Buffer)
	s.Stdout = io.MultiWriter(s.Stdout, buf)
	s.Stderr = io.MultiWriter(s.Stderr, buf)

	go func() {
		preLine := make([]byte, 0)
		for {
			if buf.Len() > 0 {
				// Read the line from the buffer until the first
				// carriage return is found.
				line, err := buf.ReadBytes('\n')
				if err == io.EOF {
					preLine = append(preLine, line...)
					continue
				}

				if _, err := fmt.Fprintf(logfile, string(append(preLine, line...))); err != nil {
					l.Errorf("[Error writing to log `%s`::%s]", c.logFile, err)
				}
				preLine = make([]byte, 0)
				continue
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return nil
}
