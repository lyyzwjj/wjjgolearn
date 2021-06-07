package _00_common

type ArrayList struct {
	BaseList
	elements []int
}

func (al *ArrayList) Get(index int) int {
	al.rangeCheck(index)
	return -1
}

func (al *ArrayList) Set(index int, element int) {
	al.rangeCheck(index)
}

func (al *ArrayList) IndexOf(element int) int {
	return -1
}

func (al *ArrayList) AddWithIndex(index int, element int) {
}

func (al *ArrayList) Remove(index int) int {
	return -1
}

func (al *ArrayList) Clear() {
}

func NewArrayListWithCapacity(capacity int) *ArrayList {
	arrayList := new(ArrayList)
	arrayList.elements = make([]int, capacity)
	// println("size: " + strconv.Itoa(len(arrayList.elements)) + "capacity: " + strconv.Itoa(cap(arrayList.elements)))
	arrayList.BaseList.AddWithIndex = arrayList.AddWithIndex
	arrayList.BaseList.Remove = arrayList.Remove
	arrayList.BaseList.IndexOf = arrayList.IndexOf
	return arrayList
}
func NewArrayList() *ArrayList {
	return NewArrayListWithCapacity(defaultCapacity)
}

/*
func (arrayList *ArrayList) ensureCapacity(capacity int) {
	oldCapacity := cap(arrayList.elements)
	if oldCapacity >= capacity {
		return
	}
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]interface{}, newCapacity)
	for i := 0; i < arrayList.size; i++ {
		newElements[i] = arrayList.elements[i]
	}
	arrayList.elements = newElements
}
func (arrayList *ArrayList) AddWithIndex(index int, element int) {
	if err := arrayList.rangeCheckForAdd(index); err != nil {
		fmt.Println(err)
		return
	}
	arrayList.ensureCapacity(arrayList.size + 1)
	for i := arrayList.size; i > index; i-- {
		arrayList.elements[i] = arrayList.elements[i-1]
	}
	arrayList.elements[index] = element
	arrayList.size++
}

func (arrayList *ArrayList) Remove(index int) (int, bool) {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return 0, false
	}
	element := arrayList.elements[index]
	for i := index; i < arrayList.size-1; i++ {
		arrayList.elements[i] = arrayList.elements[i+1]
	}
	value, _ := element.(int)
	return value, true
}
func (arrayList *ArrayList) Clear() {
	for i := 0; i < arrayList.size; i++ {
		arrayList.elements[i] = nil
	}
	arrayList.size = 0
}
func (arrayList *ArrayList) Get(index int) *int {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return nil
	}
	value, _ := arrayList.elements[index].(int)
	return &value
}
func (arrayList *ArrayList) Set(index int, element int) (int, bool) {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return 0, false
	}
	oldValue, _ := arrayList.elements[index].(int)
	arrayList.elements[index] = element
	return oldValue, true
}
func (arrayList *ArrayList) IndexOf(element int) *int {
	for i := 0; i < arrayList.size; i++ {
		if arrayList.elements[i] == element {
			return &i
		}
	}
	return nil
}
*/
