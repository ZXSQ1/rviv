package expiry

import (
	"regexp"
	"strconv"
	"time"
)

func ParseExpiry(expiryDuration string) (minTime time.Time, err error) {
	format := `([0123456789]+[yMdhms])+`
	re := regexp.MustCompile(format)

	if !re.MatchString(expiryDuration) {
		return time.Time{}, ErrInvalidDuration
	}

	resultTime := time.Now()
	numText := ""

	for _, c := range expiryDuration {
		if '0' <= c && c <= '9' {
			numText += string(c)
			continue
		}

		num, err := strconv.Atoi(numText)

		if err != nil {
			return time.Time{}, err
		}

		switch c {
		case 'y':
			resultTime = resultTime.AddDate(-num, 0, 0)
		case 'M':
			resultTime = resultTime.AddDate(0, -num, 0)
		case 'd':
			resultTime = resultTime.AddDate(0, 0, -num)
		case 'h':
			resultTime = resultTime.Add(-(time.Duration(num) * time.Hour))
		case 'm':
			resultTime = resultTime.Add(-(time.Duration(num) * time.Minute))
		case 's':
			resultTime = resultTime.Add(-(time.Duration(num) * time.Second))
		}

		numText = ""
	}

	return resultTime, nil
}
