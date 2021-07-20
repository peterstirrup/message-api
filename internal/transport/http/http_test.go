package http_test

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	transport "github.com/peterstirrup/messages/internal/transport/http"
	http_mocks "github.com/peterstirrup/messages/internal/transport/http/mock"
	openapi "github.com/peterstirrup/messages/pkg/client"
)

const clientID = int64(32)

var ctx = contextMatcher{}

type contextMatcher struct{}

func (m contextMatcher) Matches(x interface{}) bool {
	_, ok := x.(context.Context)
	return ok
}

func (m contextMatcher) String() string {
	return "a context"
}

type setupHTTPTestConfig struct {
	mockCtrl *gomock.Controller
	usecases *http_mocks.MockUseCases
	handler  http.Handler
}

func setupHTTPTest(t *testing.T) *setupHTTPTestConfig {
	ctrl := gomock.NewController(t)
	usecases := http_mocks.NewMockUseCases(ctrl)
	srv := transport.NewServer(usecases)

	handler := srv.NewHandler()

	return &setupHTTPTestConfig{
		mockCtrl: ctrl,
		usecases: usecases,
		handler:  handler,
	}
}

func teardownHTTPTest(t *testing.T, cfg *setupHTTPTestConfig) {
	cfg.mockCtrl.Finish()
}

func expectStatusCode(t *testing.T, w *httptest.ResponseRecorder, expectedCode int) {
	if w.Result().StatusCode != expectedCode {
		t.Fatalf("expected status %d, got %d with body %s", expectedCode, w.Result().StatusCode, w.Body.String())
	}
}

func expectJSONContentType(t *testing.T, w *httptest.ResponseRecorder) {
	if w.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected content type 'application/json', got %s", w.Header().Get("Content-Type"))
	}
}

func expectError(t *testing.T, w *httptest.ResponseRecorder, expectedErr openapi.Error) {
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	var httpErr openapi.Error
	if err := json.Unmarshal(body, &httpErr); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}

	if diff := cmp.Diff(httpErr, expectedErr); diff != "" {
		t.Errorf("expected %+v, got %+v", expectedErr, httpErr)
	}
}

func makeRequest(httpMethod string, path string, payload io.Reader, clientID int64) *http.Request {
	r := httptest.NewRequest(httpMethod, path, payload)
	r.Header.Set("Accept", "application/json")
	r.Header.Set("client", strconv.FormatInt(clientID, 10))
	return r
}
