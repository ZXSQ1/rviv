package expiry

import (
	"testing"
	"time"
)

func TestParseExpiry(t *testing.T) {
	if _, err := ParseExpiry("jfksdjlj3lk1lj2ljl"); err == nil {
		t.FailNow()
	}

	minTime, err := ParseExpiry("3h31m")

	if err != nil {
		t.FailNow()
	}

	if int(time.Now().Sub(minTime).Hours()) != 3 ||
		int(time.Now().Sub(minTime).Minutes()) != 211 {

		t.FailNow()
	}

	minTime, err = ParseExpiry("10y3M2d")

	if err != nil {
		t.FailNow()
	}

	if time.Now().Year()-minTime.Year() != 10 ||
		time.Now().Month()-minTime.Month() != 3 ||
		time.Now().Day()-minTime.Day() != 2 {

		t.FailNow()
	}
}
