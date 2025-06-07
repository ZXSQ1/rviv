package sftpfs

import (
	"strconv"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFtpFs struct {
	conn *sftp.Client
}

func Connect(ip string, port int, user, pass string) (*SFtpFs, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

	sshConn, err := ssh.Dial("tcp", ip+":"+strconv.Itoa(port), config)

	if err != nil {
		return nil, err
	}

	sftpConn, err := sftp.NewClient(sshConn)

	if err != nil {
		return nil, err
	}

	return &SFtpFs{
		conn: sftpConn,
	}, nil
}
