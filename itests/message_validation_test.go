package itests

import (
	. "github.com/go-clarum/go-binding/http/request"
	. "github.com/go-clarum/go-binding/http/response"
	"net/http"
	"testing"
)

// Method GET
// + single query param validation
// + URL from client
func TestGet(t *testing.T) {
	testClient.In(t).Send().
		Request(Get().QueryParam("myParam", "myValue1"))

	firstTestServer.In(t).Receive().
		Request(Get("/myApp/").QueryParam("myParam", "myValue1"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK))

	testClient.In(t).Receive().
		Name("Receive GET response").
		Response(Response(http.StatusOK))
}

// Method HEAD
// + URL overwrite
func TestHead(t *testing.T) {
	testClient.In(t).Send().
		Request(Head("myOtherApp").BaseUrl("http://localhost:8084"))

	secondTestServer.In(t).Receive().
		Request(Head("myOtherApp").BaseUrl("has no effect on server"))
	secondTestServer.In(t).Send().
		Response(Response(http.StatusOK))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK))
}

// Method POST
// + multiple query params
func TestPost(t *testing.T) {
	testClient.In(t).Send().
		Request(Post().
			QueryParam("myParam1", "myValue1").
			QueryParam("myParam2", "myValue1").
			Payload("my plain text payload"))

	firstTestServer.In(t).Receive().
		Request(Post("myApp").
			QueryParam("myParam1", "myValue1").
			QueryParam("myParam2", "myValue1").
			Payload("my plain text payload"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK))
}

// Method PUT
// + query param with multiple values
// + authorization header
// + request payload validation
func TestPut(t *testing.T) {
	testClient.In(t).Send().
		Request(Put().
			QueryParam("myParam1", "myValue1", "myValue2").
			Authorization("1234").
			Payload("my plain text payload"))

	firstTestServer.In(t).Receive().
		Request(Put("myApp").
			QueryParam("myParam1", "myValue1").
			Authorization("1234").
			Payload("my plain text payload"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusCreated))

	testClient.In(t).Receive().
		Response(Response(http.StatusCreated))
}

// Method DELETE
// + path validation
// + server ignores Authorization header
// + server ignores request payload
func TestDelete(t *testing.T) {
	testClient.In(t).Send().
		Request(Delete("my", "/", "resource", "", "1234").
			Authorization("some token which is ignored on server validation").
			Payload("payload which will be ignored"))

	firstTestServer.In(t).Receive().
		Request(Delete("myApp/my/resource/1234"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK))
}

// Method OPTIONS
// + multiple header validation server side
// + single header validation client side
// + client ignores response payload
func TestOptions(t *testing.T) {
	testClient.In(t).Send().
		Request(Options().
			Header("trace", "231561234234").
			Header("span", "33334444"))

	firstTestServer.In(t).Receive().
		Request(Options("myApp").
			Header("trace", "231561234234").
			Header("span", "33334444"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK).
			ETag("555777666").
			Payload("payload which will be ignored"))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK).
			ETag("555777666"))
}

// Method TRACE
// + response payload validation
func TestTrace(t *testing.T) {
	testClient.In(t).Send().
		Request(Trace())

	firstTestServer.In(t).Receive().
		Request(Trace("myApp"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK).
			Payload("my special response"))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK).
			Payload("my special response"))
}

// Method PATCH
func TestPatch(t *testing.T) {
	testClient.In(t).Send().
		Request(Patch())

	firstTestServer.In(t).Receive().
		Request(Patch("myApp"))
	firstTestServer.In(t).Send().
		Response(Response(http.StatusOK))

	testClient.In(t).Receive().
		Response(Response(http.StatusOK))
}
