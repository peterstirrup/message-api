package whatsapp

import (
	"net/http"

	"github.com/peterstirrup/messages/internal/messages/errors"
)

// WhatsApp is used to access WhatsApp functionality.
type WhatsApp struct {
	client *http.Client
	apiURL string
}

// New returns a WhatsApp struct with the passed in http.Client.
// If client is missing, return an error.
func New(client *http.Client, apiURL string) (*WhatsApp, error) {
	if client == nil {
		return nil, errors.HTTPClientMissing
	}

	return &WhatsApp{
		client: client,
		apiURL: apiURL,
	}, nil
}
