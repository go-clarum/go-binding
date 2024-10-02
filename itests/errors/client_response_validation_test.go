package errors

import (
	. "github.com/go-clarum/go-binding/http/request"
	. "github.com/go-clarum/go-binding/http/response"
	"net/http"
	"testing"
)

// The following tests check client receive response validation errors.

// HTTP header validation error: header missing
func TestHeaderMissingResponseValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - header <etag> missing",
	}

	e1 := errorsClient.Send().
		Request(Get().BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Get())
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK).
			ETag("132r1r312e1"))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP header validation error: header value incorrect
func TestHeaderInvalidResponseValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - header <someheader> mismatch - expected [wrongValue] but received [[someValue]]",
	}

	e1 := errorsClient.Send().
		Request(Get().BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Get())
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK).
			ETag("132r1r312e1").
			Header("someHeader", "someValue"))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK).
			ETag("132r1r312e1").
			Header("someHeader", "wrongValue"))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP plain text response payload validation error: payload missing
func TestMissingTextResponsePayloadValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - payload missing - expected [expected payload] but received no payload",
	}

	e1 := errorsClient.Send().
		Request(Get().BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Get())
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK).
			Payload("expected payload"))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP plain text response payload validation error: payload invalid
func TestWrongTextResponsePayloadValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - payload mismatch - expected [expected payload] but received [wrong payload]",
	}

	e1 := errorsClient.Send().
		Request(Get().BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Get())
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK).
			Payload("wrong payload"))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK).
			Payload("expected payload"))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}
