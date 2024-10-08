package errors

import (
	. "github.com/go-clarum/go-binding/http/request"
	. "github.com/go-clarum/go-binding/http/response"
	"net/http"
	"testing"
)

// The following tests check server validation errors.

// HTTP method validation error.
// Client sends HTTP GET & server expects POST
func TestMethodValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - method mismatch - expected [POST] but received [GET]",
	}

	e1 := errorsClient.Send().Request(Get().BaseUrl("http://localhost:8083/myApp"))

	e2 := errorsServer.Receive().Request(Post("myApp"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP status code error.
// Server receives a message to send with an invalid HTTP status code -> default error response because of the error
func TestInvalidStatusCode(t *testing.T) {
	expectedErrors := []string{
		"errorsServer: action to send is invalid - unsupported status code [99]",
		"validation error - status mismatch - expected [200] but received [500]",
	}

	e1 := errorsClient.Send().Request(Get().BaseUrl("http://localhost:8083/myApp"))

	e2 := errorsServer.Receive().Request(Get("myApp"))
	e3 := errorsServer.Send().
		Response(Response(99))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP status code validation error.
// Server responds with 400 Bad Request & client expects 200 OK
func TestStatusCodeValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - status mismatch - expected [200] but received [400]",
	}

	e1 := errorsClient.Send().Request(Get().BaseUrl("http://localhost:8083/myApp"))

	e2 := errorsServer.Receive().Request(Get("myApp"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusBadRequest))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP path validation error.
// Server responds with 404 Bad Request & client expects 200 OK
func TestPathValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - path mismatch - expected [my/resource/5433] but received [my/resource/1234]",
	}

	e1 := errorsClient.Send().
		Request(Get("my", "resource", "1234").
			BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Get("my", "resource", "5433"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusNotFound))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusNotFound))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP header validation error: multiple headers, one missing
func TestHeaderMissingValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - header <traceid> missing",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083").
			Authorization("Bearer: 123152123123"))

	e2 := errorsServer.Receive().Request(Get().
		Authorization("Bearer: 123152123123").
		Header("traceid", "777777777"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP header validation error: header value incorrect
func TestHeaderInvalidValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - header <authorization> mismatch - expected [Bearer: 234121] but received [[Bearer: 123152123123]]",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083").
			Authorization("Bearer: 123152123123"))

	e2 := errorsServer.Receive().Request(Get().
		Authorization("Bearer: 234121"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP query params validation error: query param missing
func TestQueryParamMissingValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - query param <param2> missing",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083").
			QueryParam("param1", "value1"))

	e2 := errorsServer.Receive().Request(Get().
		QueryParam("param1", "value1").
		QueryParam("param2", "value2"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP query params validation error: query param value mismatch
func TestQueryParamInvalidValueValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - query param <param2> values mismatch - expected [[value3]] but received [[value2]]",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083").
			QueryParam("param1", "value1").
			QueryParam("param2", "value2"))

	e2 := errorsServer.Receive().Request(Get().
		QueryParam("param1", "value1").
		QueryParam("param2", "value3"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP query params validation error: query param multi value mismatch
func TestQueryParamInvalidMultiValueValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - query param <param2> values mismatch - expected [[value2 value3]] but received [[value2 value4]]",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083").
			QueryParam("param1", "value1").
			QueryParam("param2", "value2", "value4"))

	e2 := errorsServer.Receive().Request(Get().
		QueryParam("param1", "value1").
		QueryParam("param2", "value2", "value3"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP plain text request payload validation error: payload missing
func TestMissingTextRequestPayloadValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - payload missing - expected [expected payload] but received no payload",
	}

	e1 := errorsClient.Send().
		Request(Post().BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().Request(Post().
		Payload("expected payload"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

// HTTP plain text request payload validation error: payload invalid
func TestWrongTextRequestPayloadValidation(t *testing.T) {
	expectedErrors := []string{
		"validation error - payload mismatch - expected [expected payload] but received [wrong payload]",
	}

	e1 := errorsClient.Send().
		Request(Post().BaseUrl("http://localhost:8083").
			Payload("wrong payload"))

	e2 := errorsServer.Receive().Request(Post().
		Payload("expected payload"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusOK))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}
