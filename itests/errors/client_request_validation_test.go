package errors

import (
	. "github.com/go-clarum/go-binding/http/request"
	"testing"
)

// The following tests check client send request validation errors.

func TestClientSendNilMessage(t *testing.T) {
	expectedErrors := []string{
		"request to send cannot be nil",
	}

	e1 := errorsClient.Send().Request(nil)

	checkErrors(t, expectedErrors, e1)
}

func TestClientSendNilUrl(t *testing.T) {
	expectedErrors := []string{
		"errorsClient: send action is invalid - missing url",
	}

	e1 := errorsClient.Send().Request(Get())

	checkErrors(t, expectedErrors, e1)
}

func TestClientSendInvalidUrl(t *testing.T) {
	expectedErrors := []string{
		"errorsClient: send action is invalid - invalid url",
	}

	e1 := errorsClient.Send().Request(Get().BaseUrl("http:/localhost:8081"))
	e2 := errorsClient.Send().Request(Get().BaseUrl("som e thi ng"))

	checkErrors(t, expectedErrors, e1)
	checkErrors(t, expectedErrors, e2)
}
