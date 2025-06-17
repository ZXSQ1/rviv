package sftpfs

import (
	"strconv"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFtpFs struct {
	conn *sftp.Client
}

func Connect(addr, user, pass string) (*SFtpFs, error) {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

	sshConn, err := ssh.Dial("tcp", addr, config)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	sftpConn, err := sftp.NewClient(sshConn)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	return &SFtpFs{
		conn: sftpConn,
	}, nil
}

func GetAddr(ip string, port int) string {
	portString := strconv.Itoa(port)
	return ip + ":" + portString
}
