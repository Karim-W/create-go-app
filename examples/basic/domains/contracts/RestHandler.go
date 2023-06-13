package contracts

type RestHandler[T any] interface {
	SetupRoutes(rg *T)
}
