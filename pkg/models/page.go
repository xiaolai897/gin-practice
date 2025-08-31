package models

type BasePage[T Model] struct {
	Total   int64 `json:"total"`
	Size    int   `json:"size"`
	Current int   `json:"current"`
	Results *[]T  `json:"results"`
}

type Page interface {
	SetTotal(total int64)
	GetTotal() int64
	GetSize() int
	GetCurrent() int
	GetResults() *[]Model
}

func (p BasePage[T]) SetTotal(total int64) {
	p.Total = total
}

func (p BasePage[T]) GetTotal() int64 {
	return p.Total
}

func (p BasePage[T]) GetSize() int {
	return p.Size
}

func (p BasePage[T]) GetCurrent() int {
	return p.Current
}

func Map[T1 any, T2 any](arr *[]T1, f func(T1) T2) *[]T2 {
	result := make([]T2, len(*arr))
	for i, elem := range *arr {
		result[i] = f(elem)
	}
	return &result
}

func (p BasePage[T]) GetResults() *[]Model {
	return Map(p.Results, func(t T) Model {
		return t
	})
}
