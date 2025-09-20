package filesystem

import (
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
