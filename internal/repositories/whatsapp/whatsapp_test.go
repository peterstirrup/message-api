package whatsapp_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peterstirrup/messages/internal/repositories/whatsapp"

	"github.com/peterstirrup/messages/internal/messages/errors"

	"github.com/peterstirrup/messages/internal/messages/entities"

	"github.com/google/go-cmp/cmp"
)

var ctx = context.Background()

type setupWhatsAppTestConfig struct {
	w   *whatsapp.WhatsApp
	srv *httptest.Server
}

func setupWhatsAppTest(t *testing.T, h http.HandlerFunc) *setupWhatsAppTestConfig {
	srv := httptest.NewServer(h)

	w, err := whatsapp.New(&http.Client{}, srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	return &setupWhatsAppTestConfig{
		w:   w,
		srv: srv,
	}
}

func teardownWhatsAppTest(cfg *setupWhatsAppTestConfig) {
	cfg.srv.Close()
}

func TestWhatsApp_GetContacts(t *testing.T) {
	t.Run("successfully returns contacts", func(t *testing.T) {
		cfg := setupWhatsAppTest(t, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
{
  "type": "contacts",
  "content": [
    {
      "jid": "447826555787@c.us",
      "name": "Danny",
      "notify": "Danny H"
    },
    {
      "jid": "96170121272@c.us",
      "name": "Karim",
      "notify": "Karim (￣^￣)ゞ"
    },
    {
      "jid": "447985413898@c.us",
      "name": "Rami",
      "notify": "Rami Kalai"
    }
  ]
}`))
		})
		defer teardownWhatsAppTest(cfg)

		got, err := cfg.w.GetContacts(ctx, 10)
		if err != nil {
			t.Fatal(err)
		}

		want := []entities.Contact{
			{ID: "447826555787@c.us", Name: "Danny"},
			{ID: "96170121272@c.us", Name: "Karim"},
			{ID: "447985413898@c.us", Name: "Rami"},
		}

		gotWant(t, got, want)
	})

	t.Run("returns err on malformed json", func(t *testing.T) {
		cfg := setupWhatsAppTest(t, func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"type`))
		})
		defer teardownWhatsAppTest(cfg)

		_, err := cfg.w.GetContacts(ctx, 10)
		if err != errors.ReceivedMalformedJSON {
			t.Fatalf("expected %s, got %s", errors.ReceivedMalformedJSON, err)
		}
	})

	t.Run("returns err when client is missing from New() call", func(t *testing.T) {
		_, err := whatsapp.New(nil, "")
		if err != errors.HTTPClientMissing {
			t.Fatalf("expected %s, got: %s", errors.HTTPClientMissing, err)
		}
	})
}

func gotWant(t *testing.T, got, want interface{}) {
	if diff := cmp.Diff(got, want); diff != "" {
		fmt.Println(diff)
		t.Fail()
	}
}
