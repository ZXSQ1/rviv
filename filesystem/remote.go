package filesystem

import (
	"strconv"
	"time"
)

// the standard object for connection information
type ConnInfo struct {
	// the address containing the ip and the port in the form: <ip>:<port>
	Addr string

	// the username
	User string

	// the password
	Pass string

	// the timeout
	Timeout time.Duration
}

const (
	// the standard default timeout for connections
	DefaultTimeout = 1 * time.Second
)

func GetAddr(ip string, port int) string {
	portString := strconv.Itoa(port)
	return ip + ":" + portString
}
