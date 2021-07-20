package conv

import "github.com/peterstirrup/messages/internal/messages/entities"

// ToTransport converts internal models into wire models types. These types are
// already configured by OpenAPI.
func ToTransport(v interface{}) interface{} {
	switch m := v.(type) {
	case []entities.Contact:
		return toTransportContacts(m)
	default:
		panic("ToTransport: unexpected type, could not convert")
	}
}
