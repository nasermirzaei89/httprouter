package httprouter

import (
	"context"
	"net/http"
	"regexp"
)

type handler struct {
	routes []route
}

// Handler interface
type Handler interface {
	http.Handler
	Handle(method, pattern string, handlerFunc http.HandlerFunc)
	Get(pattern string, handlerFunc http.HandlerFunc)
	Head(pattern string, handlerFunc http.HandlerFunc)
	Post(pattern string, handlerFunc http.HandlerFunc)
	Put(pattern string, handlerFunc http.HandlerFunc)
	Patch(pattern string, handlerFunc http.HandlerFunc)
	Delete(pattern string, handlerFunc http.HandlerFunc)
	Options(pattern string, handlerFunc http.HandlerFunc)
}

type route struct {
	method      string
	pattern     *regexp.Regexp
	handlerFunc http.HandlerFunc
}

// New return a new handler implementation
func New() Handler {
	return &handler{
		routes: []route{},
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for i := range h.routes {
		if r.Method == h.routes[i].method && h.routes[i].pattern.MatchString(r.URL.Path) {
			names := h.routes[i].pattern.SubexpNames()
			values := h.routes[i].pattern.FindAllStringSubmatch(r.URL.Path, -1)
			if len(values) > 0 {
				for i, v := range values[0] {
					if names[i] != "" {
						r = r.WithContext(context.WithValue(r.Context(), names[i], v))
					}
				}
			}
			h.routes[i].handlerFunc(w, r)
			return
		}
	}

	// not found
	http.NotFound(w, r)
}

func (h *handler) Handle(method, pattern string, handlerFunc http.HandlerFunc) {
	h.routes = append(h.routes, route{method, regexp.MustCompile(pattern), handlerFunc})
}

func (h *handler) Get(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodGet, pattern, handlerFunc)
}

func (h *handler) Head(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodHead, pattern, handlerFunc)
}

func (h *handler) Post(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodPost, pattern, handlerFunc)
}

func (h *handler) Put(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodPut, pattern, handlerFunc)
}

func (h *handler) Patch(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodPatch, pattern, handlerFunc)
}

func (h *handler) Delete(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodDelete, pattern, handlerFunc)
}

func (h *handler) Options(pattern string, handlerFunc http.HandlerFunc) {
	h.Handle(http.MethodOptions, pattern, handlerFunc)
}
