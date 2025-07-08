package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNewPath(t *testing.T) {
	t.Cleanup(func() {

		os.Remove(testFilename)
	})

	InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},
		},
	)

	testDevice := "dev"
	testFilename := "/fhjsd/fsjd[%d]"
	pathObj, err := NewPath(testDevice + "::" + testFilename)

	if err != nil {
		t.FailNow()
	}

	if pathObj.Active {
		t.FailNow()
	}

	if pathObj.Devname != testDevice {
		t.FailNow()
	}

	expectedDay := strconv.Itoa(time.Now().Day())
	expectedFilename := strings.ReplaceAll(testFilename, "%d", expectedDay)

	if len(expectedDay) == 1 {
		expectedFilename = strings.ReplaceAll(
			testFilename, "%d", "0"+expectedDay)
	}

	if pathObj.Filename != expectedFilename {
		log.Fatalln(expectedFilename, pathObj.Filename)
		t.FailNow()
	}

	if _, err := NewPath("sdfs::/sfdfs"); err == nil {
		t.FailNow()
	}

	if _, err := NewPath("/fsdfsdf"); err == nil {
		t.FailNow()
	}
}
