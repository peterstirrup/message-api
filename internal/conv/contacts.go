package conv

import (
	"github.com/peterstirrup/messages/internal/messages/entities"
	openapi "github.com/peterstirrup/messages/pkg/client"
)

func toTransportContacts(model []entities.Contact) openapi.Contacts {
	var contacts openapi.Contacts

	contacts.Contacts = make([]openapi.Contact, len(model))

	for i, contact := range model {
		contacts.Contacts[i] = openapi.Contact{
			Id:   contact.ID,
			Name: contact.Name,
		}
	}

	return contacts
}
