package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	ssc *ssh.Client
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Open(host string, port int, username string, password string) error {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	ssc, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}

	c.ssc = ssc
	return nil
}

func (c *Client) Close() error {
	if c.ssc != nil {
		return c.ssc.Close()
	}
	return nil
}

func (c *Client) NewSession() *Session {
	return NewSession(c.ssc)
}

func (c *Client) NewSftp() *Sftp {
	return NewSftp(c.ssc)
}

func (c *Client) Execute(cmd string) ([]byte, error) {
	session, err := c.ssc.NewSession()
	if err != nil {
		return nil, err
	}

	defer session.Close()
	return session.CombinedOutput(cmd)
}
