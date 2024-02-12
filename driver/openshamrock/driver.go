package openShamrockDriver

import (
	"net/url"
	"os"
)

type Transceiver interface {
	Connect(url string, authToken string) error
	SendJsonRequest(data []byte, endpoint string) ([]byte, error)
	SendFormRequest(data url.Values, endpoint string) ([]byte, error)
	SendFileRequest(data *os.File, endpoint string) ([]byte, error)
}
