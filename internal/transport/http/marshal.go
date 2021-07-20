package http

import (
	"encoding/json"
	goerr "errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	openapi "github.com/peterstirrup/messages/pkg/client"

	"github.com/peterstirrup/messages/internal/messages/errors"
)

// marshalError takes an error (from the errors package) and writes a status, and sometimes an accompanying message,
// to inform the client of the error.
func (s *Server) marshalError(err error, w http.ResponseWriter) {
	var errorMsg openapi.Error

	switch {
	case goerr.Is(err, errors.BadRequest):
		errorMsg.Status = http.StatusBadRequest
	case goerr.Is(err, errors.ClientMissing):
		errorMsg.Status = http.StatusBadRequest
		errorMsg.Detail = errors.ClientMissing.Error()
	default:
		errorMsg.Status = http.StatusInternalServerError
	}

	if errorMsg.Title == "" {
		errorMsg.Title = http.StatusText(int(errorMsg.Status))
	}

	if errorMsg.Detail == "" {
		errorMsg.Detail = err.Error()
	}

	b, err := json.Marshal(&errorMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: Log this error somewhere
	w.WriteHeader(int(errorMsg.Status))
	w.Write(b)
}

// marshalJSON takes an interface, marshals it to JSON and writes to the client.
func marshalJSON(v interface{}, w http.ResponseWriter) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(b)

	return err
}

// unmarshalJSON takes a request, reads the body and unmarshal's it into the passed in interface.
func unmarshalJSON(in io.Reader, p interface{}) error {
	b, err := ioutil.ReadAll(in)
	if err != nil {
		return fmt.Errorf("failed to read from io.Reader: %v", err)
	}

	if err := json.Unmarshal(b, p); err != nil {
		return fmt.Errorf("unable to unmarshal json: %v", err)
	}

	return nil
}
