package kernel

import (
	"context"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/router"
	"net/http"
)

type Kernel struct {
	Router  *router.GinRouter
	Modules map[string]Module
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

	notificationModule := InitNotificationModule(k, cnf)
	k.addModule(notificationModule)

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
	if c.Modules == nil {
		c.Modules = make(map[string]Module)
	}

	if c.Modules[module.Name()] != nil {
		panic("Module already exists")
	}

	c.Modules[module.Name()] = module
}

func (c *Kernel) ShutdownServer(ctx context.Context) error {
	return c.server.Shutdown(ctx)
}
