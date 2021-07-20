package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	useCases UseCases
}

// NewServer returns a pointer to a new Server instance.
func NewServer(usecases UseCases) *Server {
	return &Server{
		useCases: usecases,
	}
}

// NewHandler sets up the paths to the corresponding http handlers.
func (s *Server) NewHandler() http.Handler {
	router := mux.NewRouter()
	mw := chain(withJSONContentType)

	whatsapp := router.PathPrefix("/whatsapp").Subrouter()
	whatsapp.HandleFunc("/contacts", mw(s.requiresClient(s.getWhatsAppContactsHandler))).Methods(http.MethodGet)

	router.HandleFunc("/healthz", s.healthz).Methods(http.MethodGet)
	return router
}

func (*Server) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
