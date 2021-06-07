package _02_set

type Set interface {
	Size() int
	IsEmpty() bool
	Clear()
	Contains(element int) bool
	Add(element int)
	Remove(element int)
}
