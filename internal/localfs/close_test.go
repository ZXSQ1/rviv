package localfs

import "testing"

func TestLocalFs_Close(t *testing.T) {
	client, err := Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	if client.Close() != nil {
		t.FailNow()
	}
}
