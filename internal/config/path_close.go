package config

import "github.com/ZXSQ1/rviv/info"

func CloseConns() error {
	for dev, conn := range activeConns {
		if conn.Close() != nil {
			return info.Error("unable to close device '%s'", dev)
		}
	}

	return nil
}
