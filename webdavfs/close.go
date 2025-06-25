package webdavfs

import (
	"net"
	"net/http"
)

func (client *WebDavFs) Close() error {
	transport, ok := client.httpClient.Transport.(*http.Transport)

	if ok && !client.closed {
		transport.CloseIdleConnections()
		client.closed = true

		return nil
	}

	return net.ErrClosed
}
