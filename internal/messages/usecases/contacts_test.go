package usecases_test

import (
	"errors"
	"testing"

	"github.com/peterstirrup/messages/internal/messages/entities"
	messages_errors "github.com/peterstirrup/messages/internal/messages/errors"
)

func TestUseCases_GetWhatsAppContacts(t *testing.T) {
	t.Run("successfully returns whatsapp contacts", func(t *testing.T) {
		cfg := setupUseCasesTest(t)
		defer teardownWhatsAppTest(cfg)

		contacts := []entities.Contact{
			{
				ID:   "447826555787@c.us",
				Name: "Danny",
			},
			{
				ID:   "893487394844@c.us",
				Name: "Karim",
			},
		}

		cfg.whatsapp.EXPECT().GetContacts(ctx, testUserID).Return(contacts, nil)

		got, err := cfg.usecases.GetWhatsAppContacts(ctx, testUserID)
		if err != nil {
			t.Fatal(err)
		}

		gotWant(t, got, contacts)
	})

	t.Run("returns err when user ID is missing", func(t *testing.T) {
		cfg := setupUseCasesTest(t)
		defer teardownWhatsAppTest(cfg)

		_, err := cfg.usecases.GetWhatsAppContacts(ctx, 0)
		if err != messages_errors.ClientMissing {
			t.Fatalf("expected %s, got %s", messages_errors.ClientMissing, err)
		}
	})

	t.Run("returns err when whatsapp contacts returns an error", func(t *testing.T) {
		cfg := setupUseCasesTest(t)
		defer teardownWhatsAppTest(cfg)

		contacts := []entities.Contact{}

		cfg.whatsapp.EXPECT().GetContacts(ctx, testUserID).Return(contacts, errors.New("some err"))

		_, err := cfg.usecases.GetWhatsAppContacts(ctx, testUserID)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
