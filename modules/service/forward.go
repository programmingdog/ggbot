package service

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Forward server event to other client
type Forward struct {
	msgWebhookURL        string
	contactWebhookURL    string
	loginStateWebhookURL string
}

// NewForward ...
func newForward(msgWebhookURL, contactWebhookURL, loginStateWebhookURL string) *Forward {
	return &Forward{msgWebhookURL, contactWebhookURL, loginStateWebhookURL}
}

func (f *Forward) forward(url string, data interface{}) error {

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = http.Post(url, `application/json`, bytes.NewReader(bs))

	return err
}
