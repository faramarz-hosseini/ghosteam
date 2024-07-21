package interfaces

import "net/http"

type Base struct {
	httpClient       *http.Client
	steamAPIEndpoint string
	apiKey           string
}

func NewBase(httpClient *http.Client, steamAPIEndpoint, apiKey string) *Base {
	return &Base{
		httpClient: httpClient, steamAPIEndpoint: steamAPIEndpoint, apiKey: apiKey,
	}
}
