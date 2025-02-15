package connection

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os/exec"
	"time"
)

// pxyCmd implements the proxy.Dialer interface to support SSH ProxyCommand.
// It executes a shell command (typically ssh or nc) and uses its stdin/stdout
// as the network connection. This allows for connections through jump hosts
// or other intermediary systems.
type pxyCmd struct {
	command string
}

// Dial implements the proxy.Dialer interface. It executes the proxy command
// and returns a connection that uses the command's stdin/stdout as the
// communication channel.
func (d *pxyCmd) Dial(_, _ string) (net.Conn, error) {
	l.Debugf("Executing proxy command: %s", d.command)

	// Execute the proxy command through the shell
	// Using "sh -c" allows for complex commands with pipes and redirections
	cmd := exec.Command("sh", "-c", d.command) //nolint:gosec

	// Create pipes for stdin/stdout which will be used as the network connection
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		if err := stdin.Close(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	// Capture stderr for debugging
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	// Start the command
	if err := cmd.Start(); err != nil {
		if err := stdin.Close(); err != nil {
			return nil, err
		}
		if err := stdout.Close(); err != nil {
			return nil, err
		}
		l.Debugf("Failed to start proxy command: %v", err)
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	// Return a connection that wraps the command's stdin/stdout
	return &pxyCmdConn{
		Reader: stdout,
		Writer: stdin,
		cmd:    cmd,
	}, nil
}

// pxyCmdConn implements the net.Conn interface by wrapping the proxy command's
// stdin/stdout as a network connection. This allows the SSH client to communicate
// through the proxy command as if it were a direct TCP connection.
type pxyCmdConn struct {
	io.Reader
	io.Writer
	cmd    *exec.Cmd // The running proxy command
	closed bool      // Ensures we only close the connection once
}

// Read implements io.Reader
func (c *pxyCmdConn) Read(p []byte) (n int, err error) {
	return c.Reader.Read(p)
}

// Write implements io.Writer
func (c *pxyCmdConn) Write(p []byte) (n int, err error) {
	return c.Writer.Write(p)
}

// Close implements net.Conn Close(). It properly cleans up the proxy command
// by closing pipes and terminating the process.
func (c *pxyCmdConn) Close() error {
	if c.closed {
		return nil
	}
	c.closed = true

	l.Debug("Closing proxy command connection")

	// Close both stdin and stdout pipes
	var result error
	if err := c.Writer.(io.Closer).Close(); err != nil {
		result = err
	}

	if err := c.Reader.(io.Closer).Close(); err != nil {
		result = err
	}

	// Kill the proxy command process and wait for it to exit
	if c.cmd.Process != nil {
		if err := c.cmd.Process.Kill(); err != nil {
			result = err
		}

		if err := c.cmd.Wait(); err != nil {
			l.Debugf("Process wait error: %v", err)
		}
	}

	return result
}

// LocalAddr returns a net.Addr representing the local address of the connection.
func (c *pxyCmdConn) LocalAddr() net.Addr { return nil }

// RemoteAddr returns a net.Addr representing the remote address of the connection.
func (c *pxyCmdConn) RemoteAddr() net.Addr { return nil }

// SetDeadline sets the read and write deadlines for the connection.
func (c *pxyCmdConn) SetDeadline(_ time.Time) error { return nil }

// SetReadDeadline sets the read deadline for the connection.
func (c *pxyCmdConn) SetReadDeadline(_ time.Time) error { return nil }

// SetWriteDeadline sets the write deadline for the connection.
func (c *pxyCmdConn) SetWriteDeadline(_ time.Time) error { return nil }
