package ftpfs

import "testing"

func TestFtpFs_Close(t *testing.T) {
	server := openTestServer()
	client, err := Connect(testIp, testPort, testUser, testPass)

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		server.Shutdown()
	})

	if client.Close() != nil {
		t.FailNow()
	}
}
