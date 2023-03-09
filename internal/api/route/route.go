package route

import "net/http"

type Handler func(w http.ResponseWriter, request *http.Request)

type Route struct {
	url    string
	method string
}

func NewRoute(url string, method string) *Route {
	return &Route{
		url:    url,
		method: method,
	}
}

func (r *Route) GetData() Handler {
	return func(w http.ResponseWriter, request *http.Request) {

	}
}
