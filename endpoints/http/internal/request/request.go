package request

import (
	"fmt"
	"github.com/go-clarum/go-binding/endpoints/http/internal/message"
	"maps"
	"net/http"
	"reflect"
	"slices"
)

type HttpRequest struct {
	message.HttpMessage
	Method      string
	Url         string
	Path        []string
	QueryParams map[string][]string
}

func Get(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodGet,
		Path:   pathElements,
	}
}

func Head(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodHead,
		Path:   pathElements,
	}
}

func Post(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodPost,
		Path:   pathElements,
	}
}

func Put(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodPut,
		Path:   pathElements,
	}
}

func Delete(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodDelete,
		Path:   pathElements,
	}
}

func Options(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodOptions,
		Path:   pathElements,
	}
}

func Trace(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodTrace,
		Path:   pathElements,
	}
}

func Patch(pathElements ...string) *HttpRequest {
	return &HttpRequest{
		Method: http.MethodPatch,
		Path:   pathElements,
	}
}

// BaseUrl - While this should normally be configured only on the HTTP client,
// this is also allowed on the message so that a client can send a request to different targets.
// When used on a message passed to an HTTP server, it will do nothing.
func (request *HttpRequest) BaseUrl(baseUrl string) *HttpRequest {
	request.Url = baseUrl
	return request
}

func (request *HttpRequest) Header(key string, value string) *HttpRequest {
	request.HttpMessage.Header(key, value)
	return request
}

func (request *HttpRequest) ContentType(value string) *HttpRequest {
	request.HttpMessage.ContentType(value)
	return request
}

func (request *HttpRequest) Authorization(value string) *HttpRequest {
	request.HttpMessage.Authorization(value)
	return request
}

func (request *HttpRequest) QueryParam(key string, values ...string) *HttpRequest {
	if request.QueryParams == nil {
		request.QueryParams = make(map[string][]string)
	}

	if _, exists := request.QueryParams[key]; exists {
		for _, value := range values {
			request.QueryParams[key] = append(request.QueryParams[key], value)
		}
	} else {
		request.QueryParams[key] = values
	}

	return request
}

func (request *HttpRequest) Payload(payload string) *HttpRequest {
	request.HttpMessage.MessagePayload = payload
	return request
}

func (request *HttpRequest) Clone() *HttpRequest {
	return &HttpRequest{
		Method:      request.Method,
		Url:         request.Url,
		Path:        slices.Clone(request.Path),
		QueryParams: maps.Clone(request.QueryParams),
		HttpMessage: request.HttpMessage.Clone(),
	}
}

func (request *HttpRequest) Equals(other *HttpRequest) bool {

	if request.Method != other.Method {
		return false
	} else if request.Url != other.Url {
		return false
	} else if !slices.Equal(request.Path, other.Path) {
		return false
	} else if !maps.Equal(request.Headers, other.Headers) {
		return false
	} else if !reflect.DeepEqual(request.QueryParams, other.QueryParams) {
		return false
	} else if request.MessagePayload != other.MessagePayload {
		return false
	}
	return true
}

func (request *HttpRequest) ToString() string {
	return fmt.Sprintf(
		"["+
			"Method: %s, "+
			"BaseUrl: %s, "+
			"Path: '%s', "+
			"Headers: %s, "+
			"QueryParams: %s, "+
			"MessagePayload: %s"+
			"]",
		request.Method, request.Url, request.Path,
		request.Headers, request.QueryParams, request.MessagePayload)
}
