package filesystem

import "strconv"

type ConnInfo struct {
	Addr string
	User string
	Pass string
}

func GetAddr(ip string, port int) string {
	portString := strconv.Itoa(port)
	return ip + ":" + portString
}
