package usecases

import (
	"context"

	"github.com/peterstirrup/messages/internal/messages/entities"
)

//go:generate mockgen -destination=mock/repositories.go github.com/peterstirrup/messages/internal/messages/usecases WhatsApp

// WhatsApp is an interface representing the interactions with a WhatsApp
// data store.
type WhatsApp interface {
	GetContacts(ctx context.Context, userID int64) ([]entities.Contact, error)
}
