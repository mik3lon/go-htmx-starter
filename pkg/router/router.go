package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Router abstracts the HTTP router functionality.
type Router interface {
	Handle(method, path string, handler gin.HandlerFunc)
	WithMiddleware(middleware ...Middleware) Router
	Serve(addr string) error
	Handler() http.Handler
	ServeStatic() error
}
