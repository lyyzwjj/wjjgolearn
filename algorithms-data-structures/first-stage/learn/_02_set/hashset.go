package set

type HashSet[T comparable] struct {
	hashmap map[T]any
}

func NewHashSet[T comparable]() Set[T] {
	hashSet := &HashSet[T]{
		hashmap: make(map[T]any),
	}
	return hashSet
}

func (h *HashSet[T]) Size() int {
	return len(h.hashmap)
}

func (h *HashSet[T]) IsEmpty() bool {
	return len(h.hashmap) == 0
}

func (h *HashSet[T]) Clear() {
	h.hashmap = make(map[T]any)
}

func (h *HashSet[T]) Contains(element T) (ok bool) {
	_, ok = h.hashmap[element]
	return
}

func (h *HashSet[T]) Add(element T) {
	h.hashmap[element] = nil
}

func (h *HashSet[T]) Remove(element T) {
	delete(h.hashmap, element)
}
func (h *HashSet[T]) GetAll() (sets []T) {
	for key := range h.hashmap {
		sets = append(sets, key)
	}
	return
}
