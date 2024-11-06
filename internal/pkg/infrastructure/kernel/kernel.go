package kernel

import (
	"context"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/router"
	"net/http"
)

type Kernel struct {
	Router  *router.GinRouter
	Modules Modules
	server  *http.Server
}

// Init initializes the container with a router implementation.
func Init(cnf *config.Config) *Kernel {
	r := router.NewGinRouter()
	err := r.ServeStatic()
	if err != nil {
		panic(err)
	}

	k := &Kernel{
		Router: r,
		server: &http.Server{
			Addr:    cnf.AddressPort,
			Handler: r.Handler(),
		},
	}

	userModule := InitUserModule(k, cnf)

	k.addModule(userModule)

	k.RegisterModuleRoutes()

	return k
}

// RegisterModuleRoutes allows each module to register its routes.
func (c *Kernel) RegisterModuleRoutes() {
	for _, m := range c.Modules {
		m.RegisterRoutes(c)
	}
}

// StartServer starts the HTTP server.
func (c *Kernel) StartServer() error {
	return c.server.ListenAndServe()
}

func (c *Kernel) addModule(module Module) {
	c.Modules = append(c.Modules, module)
}

func (c *Kernel) ShutdownServer(ctx context.Context) error {
	return c.server.Shutdown(ctx)
}
