package lan

import (
	"slices"
	"strconv"
	"sync"
	"time"
)

func ScanLanAddrs(port int, timeout time.Duration) []string {
	lanAddrs := []string{}
	subnet := getActiveSubnet()
	ips := getIPs(subnet)

	if len(ips) >= 2 {
		ips = ips[1 : len(ips)-1]
	}

	var mutex = &sync.Mutex{}
	var waitGroup = &sync.WaitGroup{}

	currIpIdx := slices.Index(ips, subnet.IP.String())
	ips = slices.Delete(ips, currIpIdx, currIpIdx+1)

	for _, ip := range ips {
		waitGroup.Add(1)

		go func() {
			addr := ip + ":" + strconv.Itoa(port)

			if pingAddr(ip, timeout) == nil {
				mutex.Lock()
				lanAddrs = append(lanAddrs, addr)
				mutex.Unlock()
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	return lanAddrs
}
