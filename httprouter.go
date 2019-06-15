package httprouter

import (
	"net/http"
	"regexp"
)

type handler struct {
	routes []route
}

type Handler interface {
	http.Handler
	Get(pattern string, handlerFunc http.HandlerFunc)
	Post(pattern string, handlerFunc http.HandlerFunc)
	Put(pattern string, handlerFunc http.HandlerFunc)
	Patch(pattern string, handlerFunc http.HandlerFunc)
	Delete(pattern string, handlerFunc http.HandlerFunc)
}

type route struct {
	method      string
	pattern     *regexp.Regexp
	handlerFunc http.HandlerFunc
}

func New() Handler {
	return &handler{
		routes: []route{},
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := range h.routes {
		if r.Method == h.routes[i].method && h.routes[i].pattern.MatchString(r.URL.Path) {
			h.routes[i].handlerFunc(w, r)
			return
		}
	}

	// not found
	http.NotFound(w, r)
}

func (h *handler) Get(pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{http.MethodGet, regexp.MustCompile(pattern), handlerFunc})
}

func (h *handler) Post(pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{http.MethodPost, regexp.MustCompile(pattern), handlerFunc})
}

func (h *handler) Put(pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{http.MethodPut, regexp.MustCompile(pattern), handlerFunc})
}

func (h *handler) Patch(pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{http.MethodPatch, regexp.MustCompile(pattern), handlerFunc})
}

func (h *handler) Delete(pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{http.MethodDelete, regexp.MustCompile(pattern), handlerFunc})
}
