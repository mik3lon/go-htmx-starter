package kernel

type Modules []Module

// Module represents a module that can register routes.
type Module interface {
	RegisterRoutes(c *Kernel)
}
