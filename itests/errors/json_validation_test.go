package errors

import (
	. "github.com/go-clarum/go-binding/http/request"
	. "github.com/go-clarum/go-binding/http/response"
	"net/http"
	"testing"
)

func TestJsonServerErrorValidation(t *testing.T) {
	expectedErrors := []string{
		"[$.name] - value mismatch - expected [Bruce] but received [Bruce Wayne]",
		"[$.age] - value mismatch - expected [37] but received [38]",
		"[$.location.street] - field is missing",
		"[$.location.number] - value mismatch - expected [1007] but received [1008]",
		"[$.location.hidden] - value mismatch - expected [false] but received [true]",
	}

	e1 := errorsClient.Send().
		Request(Put().
			BaseUrl("http://localhost:8083").
			Payload("{" +
				"\"active\": true," +
				" \"name\": \"Bruce Wayne\"," +
				" \"age\": 38," +
				" \"height\": 1.879," +
				"\"location\": {" +
				"\"address\": \"Mountain Drive\"," +
				"\"number\": 1008," +
				"\"hidden\": true" +
				"}" +
				"}"))

	e2 := errorsServer.Receive().
		Json().
		Request(Put().
			Payload("{" +
				"\"active\": true," +
				" \"name\": \"Bruce\"," +
				" \"age\": 37," +
				" \"height\": 1.879," +
				"\"location\": {" +
				"\"street\": \"Mountain Drive\"," +
				"\"number\": 1007," +
				"\"hidden\": false" +
				"}" +
				"}"))
	e3 := errorsServer.Send().
		Response(Response(http.StatusInternalServerError))

	e4 := errorsClient.Receive().
		Response(Response(http.StatusInternalServerError))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}

func TestJsonClientErrorValidation(t *testing.T) {
	expectedErrors := []string{
		"[$.name] - value mismatch - expected [Bruce] but received [Bruce Wayne]",
		"[$.age] - value mismatch - expected [37] but received [38]",
		"[$.location.street] - field is missing",
		"[$.location.number] - value mismatch - expected [1007] but received [1008]",
		"[$.location.hidden] - value mismatch - expected [false] but received [true]",
	}

	e1 := errorsClient.Send().
		Request(Get().
			BaseUrl("http://localhost:8083"))

	e2 := errorsServer.Receive().
		Json().
		Request(Get())
	e3 := errorsServer.Send().
		Response(Response(http.StatusOK).
			Payload("{" +
				"\"active\": true," +
				" \"name\": \"Bruce Wayne\"," +
				" \"age\": 38," +
				" \"height\": 1.879," +
				"\"location\": {" +
				"\"address\": \"Mountain Drive\"," +
				"\"number\": 1008," +
				"\"hidden\": true" +
				"}" +
				"}"))

	e4 := errorsClient.Receive().
		Json().
		Response(Response(http.StatusOK).
			Payload("{" +
				"\"active\": true," +
				" \"name\": \"Bruce\"," +
				" \"age\": 37," +
				" \"height\": 1.879," +
				"\"location\": {" +
				"\"street\": \"Mountain Drive\"," +
				"\"number\": 1007," +
				"\"hidden\": false" +
				"}" +
				"}"))

	checkErrors(t, expectedErrors, e1, e2, e3, e4)
}
