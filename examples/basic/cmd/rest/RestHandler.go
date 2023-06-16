package rest

type RestHandler[T any] interface {
	SetupRoutes(rg *T)
}
