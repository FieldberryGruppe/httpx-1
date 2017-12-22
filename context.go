package httpx

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
)

// NewRequest creates a request using httptest.NewRequest and initializes
// its context for use with the Context function.
func NewRequest(method, address string, body io.Reader) *http.Request {
	return InitializeContext(httptest.NewRequest(method, address, body))
}

type ContextualRequest interface {
	Context() context.Context
}

func InitializeContext(request *http.Request) *http.Request {
	parent := request.Context()
	child := context.WithValue(parent, contextNamespace, make(Namespace))
	return request.WithContext(child)
}

func Context(request ContextualRequest) Namespace {
	namespace, ok := request.Context().Value(contextNamespace).(Namespace)
	if !ok {
		namespace = make(Namespace)
	}
	return namespace
}

type Namespace map[interface{}]interface{}

func (this Namespace) Int(key interface{}) int {
	value, _ := this[key].(int)
	return value
}
func (this Namespace) Int64(key interface{}) int64 {
	value, _ := this[key].(int64)
	return value
}

func (this Namespace) Uint64(key interface{}) uint64 {
	value, _ := this[key].(uint64)
	return value
}

func (this Namespace) String(key interface{}) string {
	value, _ := this[key].(string)
	return value
}

const contextNamespace = "smartystreets"
