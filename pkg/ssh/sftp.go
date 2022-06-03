package ssh

import (
	"context"
	"io"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type StatCallback = func(srcInfo os.FileInfo, srcErr error, dstInfo os.FileInfo, dstErr error)

type Sftp struct {
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

func NewSftp(sshClient *ssh.Client) *Sftp {
	return &Sftp{
		sshClient: sshClient,
	}
}

func (s *Sftp) Establish() error {
	sftpClient, err := sftp.NewClient(s.sshClient)
	if err != nil {
		return err
	}

	s.sftpClient = sftpClient
	return nil
}

func (s *Sftp) Disconnect() error {
	if s.sftpClient != nil {
		_ = s.sftpClient.Close()
	}
	return nil
}

func (s *Sftp) Pull(remoteDir, remoteFilename string, localDir, localFilename string, callbacks ...StatCallback) (int64, error) {
	remotePath := path.Join(remoteDir, remoteFilename)
	remoteFile, err := s.sftpClient.Open(remotePath)
	if err != nil {
		return 0, err
	}
	defer remoteFile.Close()

	if _, err := os.Stat(localDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(localDir, 0744)
		}
	}

	localPath := filepath.Join(localDir, localFilename)
	localFile, err := os.Create(localPath)
	if err != nil {
		return 0, err
	}
	defer localFile.Close()

	ctx, cancel := context.WithCancel(context.Background())

	wgExit := &sync.WaitGroup{}
	wgExit.Add(1)

	go func(ctx context.Context) {
		defer wgExit.Done()

		statFileInfo := func() {
			srcInfo, srcErr := s.sftpClient.Stat(remotePath)
			dstInfo, dstErr := os.Stat(localPath)
			for _, callback := range callbacks {
				callback(srcInfo, srcErr, dstInfo, dstErr)
			}
		}
		for {
			select {
			case <-ctx.Done():
				statFileInfo()
				return
			case <-time.After(time.Second * 1):
				statFileInfo()
			}
		}
	}(ctx)

	written, err := io.Copy(localFile, remoteFile)
	cancel()
	wgExit.Wait()
	return written, err
}

func (s *Sftp) Push(localDir, localFilename string, remoteDir, remoteFilename string, callbacks ...StatCallback) (int64, error) {
	localPath := filepath.Join(localDir, localFilename)
	localFile, err := os.Open(localPath)
	if err != nil {
		return 0, err
	}
	defer localFile.Close()

	remotePath := path.Join(remoteDir, remoteFilename)
	remoteFile, err := s.sftpClient.Create(remotePath)
	if err != nil {
		return 0, err
	}
	defer remoteFile.Close()

	ctx, cancel := context.WithCancel(context.Background())

	wgExit := &sync.WaitGroup{}
	wgExit.Add(1)

	go func(ctx context.Context) {
		defer wgExit.Done()

		statFileInfo := func() {
			srcInfo, srcErr := os.Stat(localPath)
			dstInfo, dstErr := s.sftpClient.Stat(remotePath)
			for _, callback := range callbacks {
				callback(srcInfo, srcErr, dstInfo, dstErr)
			}
		}
		for {
			select {
			case <-ctx.Done():
				statFileInfo()
				return
			case <-time.After(time.Second * 1):
				statFileInfo()
			}
		}
	}(ctx)

	written, err := io.Copy(remoteFile, localFile)
	cancel()
	wgExit.Wait()
	return written, err
}

func (s *Sftp) Remove(remoteDir, remoteFilename string) error {
	remotePath := path.Join(remoteDir, remoteFilename)
	return s.sftpClient.Remove(remotePath)
}
