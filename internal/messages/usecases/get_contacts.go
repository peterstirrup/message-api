package usecases

import (
	"context"

	"github.com/peterstirrup/messages/internal/messages/entities"
	"github.com/peterstirrup/messages/internal/messages/errors"
)

// GetWhatsAppContacts calls the WhatsApp repository to retrieve the contacts
// of user with userID. Errors if user ID is missing.
func (m Messages) GetWhatsAppContacts(ctx context.Context, userID int64) ([]entities.Contact, error) {
	if userID == 0 {
		return []entities.Contact{}, errors.ClientMissing
	}

	return m.WhatsApp.GetContacts(ctx, userID)
}
