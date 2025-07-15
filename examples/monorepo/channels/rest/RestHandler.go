package rest

type Controller[T any] interface {
	SetupRoutes(rg T)
}
