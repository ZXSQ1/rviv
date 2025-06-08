package ftpfs

import (
	"log"
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
		"docker", "run", "-p", strconv.Itoa(testPort)+":21",
		"-p", "31000-31009:31000-31009",
		"-e", "FTP_USER_NAME="+testUser,
		"-e", "FTP_USER_PASS="+testPass,
		"-e", "PUBLICHOST="+testIp,
		"-e", "FTP_PASSIVE_PORTS=31000:31009",
		"stilliard/pure-ftpd:hardened",
	)

	if cmd.Start() != nil {
		log.Fatalln("failed to run FTP server")
	}

	return cmd.Process
}
