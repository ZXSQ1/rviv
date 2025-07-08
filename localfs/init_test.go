package localfs

import (
	"os"
	"testing"
)

var testPrefix = os.TempDir()

func TestInit(t *testing.T) {
	client, err := Init(testPrefix)

	if err != nil {
		t.FailNow()
	}

	if client == nil {
		t.FailNow()
	}
}
