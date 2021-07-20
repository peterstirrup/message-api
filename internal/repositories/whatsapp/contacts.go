package whatsapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peterstirrup/messages/internal/messages/errors"

	"github.com/peterstirrup/messages/internal/messages/entities"

	"go.opencensus.io/trace"
)

type contacts struct {
	Type    string `json:"type"`
	Content []struct {
		JID    string `json:"jid"`
		Name   string `json:"name"`
		Notify string `json:"notify"`
	} `json:"content"`
}

// GetContacts returns a contact list for the user with passed in userID.
func (w *WhatsApp) GetContacts(ctx context.Context, userID int64) ([]entities.Contact, error) {
	ctx, span := trace.StartSpan(ctx, "whatsapp.GetContacts")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%d/contacts", w.apiURL, 10), nil)
	if err != nil {
		return []entities.Contact{}, errors.FailedToDoHTTPRequest
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return []entities.Contact{}, errors.FailedToDoHTTPRequest
	}
	defer resp.Body.Close()

	var rawContacts contacts
	if err = json.NewDecoder(resp.Body).Decode(&rawContacts); err != nil {
		return []entities.Contact{}, errors.ReceivedMalformedJSON
	}

	contacts := make([]entities.Contact, len(rawContacts.Content))
	for i, c := range rawContacts.Content {
		contacts[i] = entities.Contact{
			ID:   c.JID,
			Name: c.Name,
		}
	}

	return contacts, nil
}
