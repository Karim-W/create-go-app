package listable

type CountedList[T any] struct {
	Data  []T `json:"data"`
	Count int `json:"count"`
}
