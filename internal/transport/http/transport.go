package http

import (
	"net/http"
	"strconv"

	"github.com/peterstirrup/messages/internal/messages/errors"
)

/*
	Would usually have this as a package that can be pulled in across
	different repositories, since this functionality would be share across
	many APIs.
*/

// middleware is a function type alias representing a single item in a
// middleware stack.
type middleware func(http.HandlerFunc) http.HandlerFunc

// requiresClientHandler is a handler that requires the client ID to be passed in.
type requiresClientHandler func(w http.ResponseWriter, req *http.Request, clientID int64)

// chain returns a function taking a HandlerFunc. This function iterates
// through a variadic list of HandlerFunc, representing the middleware stack,
// passing to each a reference to the previous function.
func chain(fns ...middleware) middleware {
	return func(route http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler := route
			for _, fn := range fns {
				handler = fn(handler)
			}

			handler(w, r)
		}
	}
}

const (
	applicationJSON = "application/json"
)

var (
	headerContentType = http.CanonicalHeaderKey("content-type")
)

// withJSONContentType set the Content-Type HTTP header on the response with a
// value indicating the payload is JSON.
func withJSONContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentType, applicationJSON)
		next.ServeHTTP(w, r)
	}
}

type (
	// clientFunc is a function that requires a client to run.
	clientFunc func(w http.ResponseWriter, req *http.Request, clientID int64)
)

// requiresClient is a middleware which wraps a HTTP handler func.
// Gets the clientID from the "client" header in the request.
// Does NOT validate the ID, or where the request came from.
func (s *Server) requiresClient(f clientFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		client := req.Header.Get("client")
		if client == "" {
			s.marshalError(errors.ClientMissing, w)
		}

		clientID, err := strconv.ParseInt(client, 10, 64)
		if err != nil {
			s.marshalError(errors.BadRequest, w)
			return
		}

		f(w, req, clientID)
	}
}
