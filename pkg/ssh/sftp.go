package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"path/filepath"
)

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

func (s *Sftp) Pull(remoteDir, remoteFilename string, localDir, localFilename string) (int64, error) {
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

	written, err := io.Copy(localFile, remoteFile)
	if err != nil {
		return 0, err
	}

	return written, nil
}

func (s *Sftp) Push(localDir, localFilename string, remoteDir, remoteFilename string) (int64, error) {
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

	written, err := io.Copy(remoteFile, localFile)
	if err != nil {
		return 0, err
	}

	return written, nil
}

func (s *Sftp) Remove(remoteDir, remoteFilename string) error {
	remotePath := path.Join(remoteDir, remoteFilename)
	return s.sftpClient.Remove(remotePath)
}
