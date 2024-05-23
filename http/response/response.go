package response

import (
	"fmt"
	"github.com/go-clarum/go-binding/http/internal/message"
	"maps"
	"strconv"
)

type HttpResponse struct {
	message.HttpMessage
	StatusCode int
}

func Response(statusCode int) *HttpResponse {
	return &HttpResponse{
		StatusCode: statusCode,
	}
}

func (response *HttpResponse) Header(key string, value string) *HttpResponse {
	response.HttpMessage.Header(key, value)
	return response
}

func (response *HttpResponse) ContentType(value string) *HttpResponse {
	response.HttpMessage.ContentType(value)
	return response
}

func (response *HttpResponse) ETag(value string) *HttpResponse {
	response.HttpMessage.ETag(value)
	return response
}

func (response *HttpResponse) Payload(payload string) *HttpResponse {
	response.HttpMessage.MessagePayload = payload
	return response
}

func (response *HttpResponse) Clone() *HttpResponse {
	return &HttpResponse{
		StatusCode:  response.StatusCode,
		HttpMessage: response.HttpMessage.Clone(),
	}
}

func (response *HttpResponse) Equals(other *HttpResponse) bool {
	if response.StatusCode != other.StatusCode {
		return false
	} else if !maps.Equal(response.Headers, other.Headers) {
		return false
	} else if response.MessagePayload != other.MessagePayload {
		return false
	}
	return true
}

func (response *HttpResponse) ToString() string {
	statusCodeText := "none"
	if response.StatusCode > 0 {
		statusCodeText = strconv.Itoa(response.StatusCode)
	}

	return fmt.Sprintf(
		"["+
			"StatusCode: %s, "+
			"Headers: %s, "+
			"Payload: %s"+
			"]",
		statusCodeText, response.Headers, response.MessagePayload)
}
