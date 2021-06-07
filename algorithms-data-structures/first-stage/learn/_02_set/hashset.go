package _02_set

type HashSet struct {
	hashmap map[int]interface{}
}

func NewHashSet() *HashSet {
	hashSet := &HashSet{
		hashmap: make(map[int]interface{}),
	}
	return hashSet
}

func (h *HashSet) Size() int {
	return len(h.hashmap)
}

func (h *HashSet) IsEmpty() bool {
	return len(h.hashmap) == 0
}

func (h *HashSet) Clear() {
	h.hashmap = make(map[int]interface{})
}

func (h *HashSet) Contains(element int) (ok bool) {
	_, ok = h.hashmap[element]
	return
}

func (h *HashSet) Add(element int) {
	h.hashmap[element] = nil
}

func (h *HashSet) Remove(element int) {
	delete(h.hashmap, element)
}
