package http

import (
	"context"

	"github.com/peterstirrup/messages/internal/messages/entities"
)

//go:generate mockgen -destination=mock/usecase.go github.com/peterstirrup/messages/internal/transport/http UseCases

type UseCases interface {
	GetWhatsAppContacts(ctx context.Context, userID int64) ([]entities.Contact, error)
}
