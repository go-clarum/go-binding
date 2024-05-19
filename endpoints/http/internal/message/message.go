package message

import (
	"github.com/go-clarum/go-binding/endpoints/http/internal/constants"
	"maps"
)

type HttpMessage struct {
	Headers        map[string]string
	MessagePayload string
}

func (message *HttpMessage) Header(key string, value string) *HttpMessage {
	if message.Headers == nil {
		message.Headers = make(map[string]string)
	}

	message.Headers[key] = value
	return message
}

func (message *HttpMessage) ContentType(value string) *HttpMessage {
	return message.Header(constants.ContentTypeHeaderName, value)
}

func (message *HttpMessage) Authorization(value string) *HttpMessage {
	return message.Header(constants.AuthorizationHeaderName, value)
}

func (message *HttpMessage) ETag(value string) *HttpMessage {
	return message.Header(constants.ETagHeaderName, value)
}

func (message *HttpMessage) Payload(payload string) *HttpMessage {
	message.MessagePayload = payload
	return message
}

func (message *HttpMessage) Clone() HttpMessage {
	return HttpMessage{
		Headers:        maps.Clone(message.Headers),
		MessagePayload: message.MessagePayload,
	}
}
