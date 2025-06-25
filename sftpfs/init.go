package sftpfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFtpFs struct {
	conn     *sftp.Client
	connInfo *filesystem.ConnInfo
}

func Connect(connInfo *filesystem.ConnInfo) (*SFtpFs, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            connInfo.User,
		Timeout:         connInfo.Timeout,
		Auth: []ssh.AuthMethod{
			ssh.Password(connInfo.Pass),
		},
	}

	sshConn, err := ssh.Dial("tcp", connInfo.Addr, config)

	if err != nil {
		return nil, err
	}

	sftpConn, err := sftp.NewClient(sshConn)

	if err != nil {
		return nil, err
	}

	return &SFtpFs{
		conn:     sftpConn,
		connInfo: connInfo,
	}, nil
}
