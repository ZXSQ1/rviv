package webdavfs

import "net/http"

func (client *WebDavFs) Close() error {
	if transport, ok := client.httpClient.Transport.(*http.Transport); ok {
		transport.CloseIdleConnections()
	}

	return nil
}
