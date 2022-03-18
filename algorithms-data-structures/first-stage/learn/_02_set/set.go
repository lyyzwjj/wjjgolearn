package set

type Set[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Contains(element T) bool
	Add(element T)
	Remove(element T)
	GetAll() (sets []T)
}
