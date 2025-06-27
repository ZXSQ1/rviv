package localfs

import "testing"

func TestLocalFs_Close(t *testing.T) {
	client := Init("/")

	if client.Close() != nil {
		t.FailNow()
	}
}
