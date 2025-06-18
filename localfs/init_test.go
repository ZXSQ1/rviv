package localfs

import "testing"

func TestInit(t *testing.T) {
	client := Init()

	if client == nil {
		t.FailNow()
	}
}
