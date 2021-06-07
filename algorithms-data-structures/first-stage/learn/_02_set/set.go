package set

type Set interface {
	Size() int
	IsEmpty() bool
	Clear()
	Contains(element interface{}) bool
	Add(element interface{})
	Remove(element interface{})
}
