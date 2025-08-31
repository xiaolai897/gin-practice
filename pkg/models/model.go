package models

type Model interface {
	Get() int64
}

func (b Base) Get() int64 {
	return b.ID
}
