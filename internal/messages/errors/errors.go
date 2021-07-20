package errors

import "errors"

var (
	// BadRequest is a catch all for bad requests.
	BadRequest = errors.New("bad request")
	// ClientMissing occurs when the user ID is missing from a call.
	ClientMissing = errors.New("client (user id) missing")
	// FailedToDoHTTPRequest is a catch all for failing to do a HTTP call.
	FailedToDoHTTPRequest = errors.New("http request failed")
	// HTTPClientMissing occurs when a http client is missing from a call when
	// it is expected.
	HTTPClientMissing = errors.New("http client missing")
	// ReceivedMalformedJSON occurs when the JSON returned cannot be
	// unmarshalled.
	ReceivedMalformedJSON = errors.New("json malformed")
)
