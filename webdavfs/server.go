package webdavfs

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
		"docker", "run", "-p", strconv.Itoa(testPort)+":80",
		"-e", "USERNAME="+testUser,
		"-e", "PASSWORD="+testPass,
		"bytemark/webdav",
	)

	cmd.Start()

	return cmd.Process
}
