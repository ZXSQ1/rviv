package ftpfs

import (
	"strconv"

	"github.com/jlaffaye/ftp"
)

type Client struct {
	conn *ftp.ServerConn
}

func Connect(ip string, port int, user, pass string) (*Client, error) {
	portString := strconv.Itoa(port)
	addr := ip + ":" + portString
	conn, err := ftp.Dial(addr)

	if err != nil {
		return nil, err
	}

	if conn.Login(user, pass) != nil {
		return nil, err
	}

	return &Client{conn}, nil
}
