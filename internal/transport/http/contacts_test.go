package http_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"

	openapi "github.com/peterstirrup/messages/pkg/client"

	"github.com/peterstirrup/messages/internal/messages/entities"
)

func TestHTTP_GetWhatsAppContacts(t *testing.T) {
	t.Run("successfully returns contacts", func(t *testing.T) {
		cfg := setupHTTPTest(t)
		defer teardownHTTPTest(t, cfg)

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

		expectedResp := openapi.Contacts{
			Contacts: []openapi.Contact{
				{
					Id:   "447826555787@c.us",
					Name: "Danny",
				},
				{
					Id:   "893487394844@c.us",
					Name: "Karim",
				},
			},
		}

		cfg.usecases.EXPECT().GetWhatsAppContacts(ctx, clientID).Return(contacts, nil)

		r := makeRequest(http.MethodGet, "/whatsapp/contacts", nil, clientID)
		w := httptest.NewRecorder()
		cfg.handler.ServeHTTP(w, r)

		expectStatusCode(t, w, http.StatusOK)
		expectJSONContentType(t, w)
		expectContacts(t, w, expectedResp)
	})
}

func expectContacts(t *testing.T, w *httptest.ResponseRecorder, expectedContacts openapi.Contacts) {
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	var contacts openapi.Contacts
	if err := json.Unmarshal(body, &contacts); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	if diff := cmp.Diff(contacts, expectedContacts); diff != "" {
		t.Errorf("expected %+v, got %+v", expectedContacts, contacts)
	}
}
