package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	router := router{handlers: make(map[string]HandlerFunc)}
	return &router
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	print("********", key)
	for key, value := range r.handlers {
		print(key, value, "\n")
	}
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND", c.Path)
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s -%s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}
