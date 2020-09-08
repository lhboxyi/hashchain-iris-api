package util

type ArrayListContainArrayList struct {
	elements []ArrayList
	size     int
}



func NewArrayListContainArrayList(values ...ArrayList) *ArrayListContainArrayList {
	list := &ArrayListContainArrayList{}
	list.elements = make([]ArrayList, 10)
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *ArrayListContainArrayList) Add(values ...ArrayList) {
	if list.size+len(values) >= len(list.elements)-1 {
		newElements := make([]ArrayList, list.size+len(values)+1)
		copy(newElements, list.elements)
		list.elements = newElements
	}

	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}

}

func (list *ArrayListContainArrayList) Remove(index int) (ArrayList,error) {
	if index < 0 || index >= list.size {
		return ArrayList{},nil
	}

	curEle := list.elements[index]
	list.elements[index] = ArrayList{}
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	return curEle,nil
}

func (list *ArrayListContainArrayList) Get(index int) (ArrayList,error) {
	if index < 0 || index >= list.size {
		return ArrayList{},nil
	}
	return list.elements[index],nil
}

func (list *ArrayListContainArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayListContainArrayList) Size() int {
	return list.size
}
/*func (list *ArrayListContainArrayList) Contains(value ArrayListContainArrayList) bool {
	for _, curValue := range list.elements {
		if curValue == value {
			return true
		}
	}

	return false
}*/
