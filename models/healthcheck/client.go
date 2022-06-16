package healthcheck

import "net/http"

type Client struct{}

func (c Client) Healthcheck() (*http.Response, error) {
	return http.Get("http://localhost:3000/health")
}
