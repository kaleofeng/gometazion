package ssh

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/sftp"
	"github.com/povsister/scp"
	"golang.org/x/crypto/ssh"
)

type Scp struct {
	sshClient  *ssh.Client
	sftpClient *sftp.Client
	scpClient  *scp.Client
}

func NewScp(sshClient *ssh.Client) *Scp {
	return &Scp{
		sshClient: sshClient,
	}
}

func (s *Scp) Establish() error {
	sftpClient, err := sftp.NewClient(s.sshClient)
	if err != nil {
		return err
	}

	scpClient, err := scp.NewClientFromExistingSSH(s.sshClient, &scp.ClientOption{})
	if err != nil {
		return err
	}

	s.sftpClient = sftpClient
	s.scpClient = scpClient
	return nil
}

func (s *Scp) Disconnect() error {
	if s.sftpClient != nil {
		_ = s.sftpClient.Close()
	}
	if s.scpClient != nil {
		_ = s.scpClient.Close()
	}
	return nil
}

func (s *Scp) Pull(remoteDir, remoteFilename string, localDir, localFilename string, callbacks ...StatCallback) (int64, error) {
	remotePath := path.Join(remoteDir, remoteFilename)
	remoteStat, err := s.sftpClient.Stat(remotePath)
	if err != nil {
		return 0, err
	}

	localPath := filepath.Join(localDir, localFilename)
	_ = os.MkdirAll(localDir, 0755)
	_ = os.Remove(localPath)

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

	err = s.scpClient.CopyFileFromRemote(remotePath, localPath, &scp.FileTransferOption{
		Context:      ctx,
		PreserveProp: true,
	})
	cancel()
	wgExit.Wait()
	return remoteStat.Size(), err
}

func (s *Scp) Push(localDir, localFilename string, remoteDir, remoteFilename string, callbacks ...StatCallback) (int64, error) {
	localPath := filepath.Join(localDir, localFilename)
	localStat, err := os.Stat(localPath)
	if err != nil {
		return 0, err
	}

	remotePath := path.Join(remoteDir, remoteFilename)
	_ = s.sftpClient.MkdirAll(remoteDir)
	_ = s.sftpClient.Remove(remotePath)

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

	err = s.scpClient.CopyFileToRemote(localPath, remotePath, &scp.FileTransferOption{
		Context:      ctx,
		PreserveProp: true,
	})
	cancel()
	wgExit.Wait()
	return localStat.Size(), err
}

func (s *Scp) Remove(remoteDir, remoteFilename string) error {
	remotePath := path.Join(remoteDir, remoteFilename)
	return s.sftpClient.Remove(remotePath)
}
