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
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
	"golang.org/x/term"
)

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

	// Create new ssh.ClientConfig{}
	config := &ssh.ClientConfig{
		User:    user,
		Auth:    []ssh.AuthMethod{ssh.Password(pass)},
		Timeout: time.Duration(timeout) * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
		Config: ssh.Config{
			KeyExchanges: []string{
				"diffie-hellman-group1-sha1",
				"diffie-hellman-group14-sha1",
				"diffie-hellman-group14-sha256",
				"ecdh-sha2-nistp256",
				"ecdh-sha2-nistp384",
				"ecdh-sha2-nistp521",
				"curve25519-sha256@libssh.org",
				"curve25519-sha256",
			},
			Ciphers: []string{
				"aes128-gcm@openssh.com",
				"chacha20-poly1305@openssh.com",
				"aes128-ctr",
				"aes192-ctr",
				"aes256-ctr",
				"aes128-cbc",
				"aes192-cbc",
				"aes256-cbc",
			},
		},
	}

	// Commenting out because we don't use this... But who knows, so I'll leave it here.
	//if hideBanner {
	//	config.BannerCallback = c.bannerCallback()
	//}

	if c.ProxyDialer == nil {
		c.ProxyDialer = proxy.Direct
	}

	// Dial to host:port
	netConn, err := c.ProxyDialer.Dial("tcp", host)
	if err != nil {
		return c, fmt.Errorf("ProxyDialer.Dial(%w)", err)
	}

	// Create new ssh connect
	sshCon, channel, req, err := ssh.NewClientConn(netConn, host, config)
	if err != nil {
		return c, fmt.Errorf("ssh.NewClientConn(%w)", err)
	}

	// Create *ssh.Client
	c.Client = ssh.NewClient(sshCon, channel, req)
	return c, nil
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
				logrus.Errorf("[session.Close::%s]", err)
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
					logrus.Errorf("[term.GetSize::%s]", err)
				}

				if err := session.WindowChange(h, w); err != nil {
					logrus.Errorf("[session.WindowChange::%s]", err)
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
		logrus.Errorf("[user.Current::%s]", err)
		return
	}

	fullPath := fmt.Sprintf("%s/.log/dmux/ssh", prof.HomeDir)
	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
			logrus.Errorf("[SetLog.Mkdir::%s]", err)
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
			logrus.Errorf("[Error setting up log path for session::%s]", err)
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
					logrus.Errorf("[Error writing to log `%s`::%s]", c.logFile, err)
				}
				preLine = make([]byte, 0)
				continue
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return nil
}
