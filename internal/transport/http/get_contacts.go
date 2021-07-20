package http

import (
	"net/http"

	"github.com/peterstirrup/messages/internal/conv"
)

// getWhatsAppContactsHandler returns a clients WhatsApp contact list.
func (s *Server) getWhatsAppContactsHandler(w http.ResponseWriter, req *http.Request, clientID int64) {
	contacts, err := s.useCases.GetWhatsAppContacts(req.Context(), clientID)
	if err != nil {
		s.marshalError(err, w)
		return
	}

	if err := marshalJSON(conv.ToTransport(contacts), w); err != nil {
		s.marshalError(err, w)
	}
}
