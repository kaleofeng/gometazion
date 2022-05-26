package ssh

import (
	"fmt"
	"io"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type Session struct {
	sshClient      *ssh.Client
	sshSession     *ssh.Session
	inputPipe      io.WriteCloser
	outputBuffer   *OutputBuffer
	outputDuration time.Duration
	outputWriter   io.Writer
	leaveChan      chan bool
	wgExit         sync.WaitGroup
}

func NewSession(sshClient *ssh.Client) *Session {
	return &Session{
		sshClient:      sshClient,
		outputBuffer:   NewOutputBuffer(),
		outputDuration: time.Millisecond * 1000,
		leaveChan:      make(chan bool),
	}
}

func (s *Session) Establish() error {
	sshSession, err := s.sshClient.NewSession()
	if err != nil {
		return err
	}

	stdinPipe, err := sshSession.StdinPipe()
	if err != nil {
		return err
	}

	sshSession.Stdout = s.outputBuffer
	sshSession.Stderr = s.outputBuffer

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := sshSession.RequestPty("vt220", 32, 160, modes); err != nil {
		return err
	}

	if err := sshSession.Shell(); err != nil {
		return err
	}

	// Echo off mode above doesn't work for root for some reason
	stdinPipe.Write([]byte("stty -echo\n"))

	s.inputPipe = stdinPipe
	s.sshSession = sshSession
	return nil
}

func (s *Session) Disconnect() error {
	if s.sshSession != nil {
		_ = s.sshSession.Close()
	}

	s.wgExit.Wait()
	return nil
}

func (s *Session) SetWriteCallback(writer io.Writer) {
	s.outputWriter = writer
}

func (s *Session) Execute(cmd string) error {
	text := fmt.Sprintf("%s\n", cmd)
	if _, err := s.inputPipe.Write([]byte(text)); err != nil {
		return err
	}
	return nil
}

func (s *Session) Watch() error {
	s.wgExit.Add(2)
	go s.wait()
	go s.output()
	return nil
}

func (s *Session) wait() {
	defer s.wgExit.Done()

	if err := s.sshSession.Wait(); err != nil {
		s.leaveChan <- true
	}
}

func (s *Session) output() {
	defer s.wgExit.Done()

	tick := time.NewTicker(s.outputDuration)
	defer tick.Stop()

	for {
		select {
		case <-s.leaveChan:
			return
		case <-tick.C:
			data := s.outputBuffer.Shift()
			_, _ = s.outputWriter.Write(data)
		}
	}
}
