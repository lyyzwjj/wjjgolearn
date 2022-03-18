package common

type Equal[T any] interface {
	Equals(t T) bool
}
type Hash[T any] interface {
	HashCode() int
}
type Object[T any] interface {
	Equal[T]
	Hash[T]
}

type Comparable[T any] interface {
	Equal[T]
	CompareTo(t T) int
}

type Comparator func(a, b interface{}) int

// IntComparator 小顶堆
func IntComparator(a, b interface{}) int {
	ia := a.(int)
	ib := b.(int)
	switch {
	case ib > ia:
		return 1
	case ib < ia:
		return -1
	default:
		return 0
	}
}

// IntRevertComparator 大顶堆
func IntRevertComparator(a, b interface{}) int {
	ia := a.(int)
	ib := b.(int)
	switch {
	case ib > ia:
		return -1
	case ib < ia:
		return 1
	default:
		return 0
	}
}

func StringComparator(a, b interface{}) int {
	ia := a.(string)
	ib := b.(string)
	switch {
	case ib > ia:
		return 1
	case ib < ia:
		return -1
	default:
		return 0
	}
}

func floatComparator(a, b interface{}) int {
	ia := a.(float64)
	ib := b.(float64)
	switch {
	case ia > ib:
		return 1
	case ia < ib:
		return -1
	default:
		return 0
	}
}
