package expiry

import (
	"math"
	"time"
)

const (
	monthAvg = 30.4375
)

var (
	durationDay   = time.Duration(time.Hour) * 24
	durationYear  = durationMonth * 12
	durationMonth = time.Duration(
		math.Round(float64(durationDay) * monthAvg),
	)
)
