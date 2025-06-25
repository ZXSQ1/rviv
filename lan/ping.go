package lan

import (
	"fmt"
	"time"

	ping "github.com/go-ping/ping"
)

func pingAddr(ip string, timeout time.Duration) error {
	pinger, err := ping.NewPinger(ip)

	if err != nil {
		return err
	}

	pinger.Count = 3
	pinger.Timeout = timeout

	if err = pinger.Run(); err != nil {
		return err
	}

	stats := pinger.Statistics()

	if stats.PacketsRecv == 0 {
		return fmt.Errorf("dialed host not alive")
	}

	return nil
}
