package listable

type QueryList[T any] struct {
	Data  []T `json:"data"`
	Count int `json:"count"`
}
