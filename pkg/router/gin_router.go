package router

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GinRouter is an adapter for the gin router.
type GinRouter struct {
	Router     *gin.Engine
	middleware []Middleware
}

func (g *GinRouter) ServeStatic() error {
	absPath, err := filepath.Abs("./assets")
	if err != nil {
		return err
	}

	fmt.Println("Serving static files from:", absPath)

	g.Router.Static("/assets", absPath)
	return nil
}

func (g *GinRouter) Handler() http.Handler {
	return g.Router
}

// NewGinRouter creates a new instance of GinRouter.
func NewGinRouter() Router {
	return &GinRouter{
		Router: gin.Default(),
	}
}

// WithMiddleware adds middleware to the router and returns the updated router.
func (g *GinRouter) WithMiddleware(middleware ...Middleware) Router {
	g.middleware = append(g.middleware, middleware...)
	return g
}

// Handle registers a new route and applies the middleware to the handler.
func (g *GinRouter) Handle(method, path string, handler gin.HandlerFunc) {
	// Wrap handler with middleware
	finalHandler := handler
	for _, m := range g.middleware {
		finalHandler = func(c *gin.Context) {
			m(func(w http.ResponseWriter, r *http.Request) {
				c.Request = r
				handler(c)
			})(c.Writer, c.Request)
		}
	}

	// Register the route with gin
	switch method {
	case http.MethodGet:
		g.Router.GET(path, finalHandler)
	case http.MethodPost:
		g.Router.POST(path, finalHandler)
	}
}

// Serve starts the Gin server at the given address.
func (g *GinRouter) Serve(addr string) error {
	return g.Router.Run(addr)
}
