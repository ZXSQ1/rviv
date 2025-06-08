package sftpfs

import (
	"os"
	"os/exec"
	"strconv"
)

var (
	testIp   = "0.0.0.0"
	testPort = 3000
	testUser = "test"
	testPass = "test"
)

func openTestServer() *os.Process {
	cmd := exec.Command(
		"docker", "run", "-p", strconv.Itoa(testPort)+":22",
		"-e", "SFTP_USERS="+testUser+":"+testPass+":1001",
		"atmoz/sftp",
	)

	cmd.Start()

	return cmd.Process
}
