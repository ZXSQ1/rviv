package ftpfs

import "testing"

func TestFtpFs_Connect(t *testing.T) {
	server := openTestServer()

	t.Cleanup(func() {
		server.Shutdown()
	})

	if _, err := Connect(testIp, testPort, testUser, testPass); err != nil {
		t.FailNow()
	}

	if _, err := Connect(testIp, 932, "190", "1902"); err == nil {
		t.FailNow()
	}
}
