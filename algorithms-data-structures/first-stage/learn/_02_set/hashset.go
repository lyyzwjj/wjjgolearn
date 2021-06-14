package set

type HashSet struct {
	hashmap map[interface{}]interface{}
}

func NewHashSet() Set {
	hashSet := &HashSet{
		hashmap: make(map[interface{}]interface{}),
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
	h.hashmap = make(map[interface{}]interface{})
}

func (h *HashSet) Contains(element interface{}) (ok bool) {
	_, ok = h.hashmap[element]
	return
}

func (h *HashSet) Add(element interface{}) {
	h.hashmap[element] = nil
}

func (h *HashSet) Remove(element interface{}) {
	delete(h.hashmap, element)
}
func (h *HashSet) GetAll() (sets []interface{}) {
	for key := range h.hashmap {
		sets = append(sets, key)
	}
	return
}
