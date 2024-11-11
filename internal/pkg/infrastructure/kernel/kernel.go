package kernel

import (
	"context"
	"github.com/mik3lon/go-htmx-starter/pkg/bus/command"
	"github.com/mik3lon/go-htmx-starter/pkg/bus/query"
	"github.com/mik3lon/go-htmx-starter/pkg/config"
	"github.com/mik3lon/go-htmx-starter/pkg/router"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type Kernel struct {
	Router     *router.GinRouter
	Modules    map[string]Module
	server     *http.Server
	CommandBus *command.CommandBus
	QueryBus   *query.QueryBus
}

// Init initializes the container with a router implementation.
func Init(cnf *config.Config) *Kernel {
	r := router.NewGinRouter()
	err := r.ServeStatic()
	if err != nil {
		panic(err)
	}

	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	k := &Kernel{
		Router: r,
		server: &http.Server{
			Addr:    cnf.AddressPort,
			Handler: r.Handler(),
		},
		CommandBus: command.InitCommandBus(l),
		QueryBus:   query.InitQueryBus(l),
	}

	userModule := InitUserModule(k, cnf)
	k.addModule(userModule)

	notificationModule := InitNotificationModule(k, cnf)
	k.addModule(notificationModule)

	k.RegisterModuleRoutes()

	return k
}

// RegisterModuleRoutes allows each module to register its routes.
func (k *Kernel) RegisterModuleRoutes() {
	for _, m := range k.Modules {
		m.RegisterRoutes(k)
	}
}

// StartServer starts the HTTP server.
func (k *Kernel) StartServer() error {
	return k.server.ListenAndServe()
}

func (k *Kernel) addModule(module Module) {
	if k.Modules == nil {
		k.Modules = make(map[string]Module)
	}

	if k.Modules[module.Name()] != nil {
		panic("Module already exists")
	}
	k.Modules[module.Name()] = module

	for c, ch := range module.Commands() {
		err := k.CommandBus.RegisterCommand(c, ch)
		if err != nil {
			panic(err)
		}
	}

	for q, ch := range module.Queries() {
		err := k.QueryBus.RegisterQuery(q, ch)
		if err != nil {
			panic(err)
		}
	}
}

func (k *Kernel) ShutdownServer(ctx context.Context) error {
	return k.server.Shutdown(ctx)
}
