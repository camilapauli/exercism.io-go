package listops

//IntList ...
type IntList []int

type unaryFunc func(x int) int
type predFunc func(n int) bool
type binFunc func(x, y int) int

//Length ...
func (il IntList) Length() int {
	return len(il)
}

//Reverse ...
func (il IntList) Reverse() IntList {
	newList := IntList{}
	lenght := len(il)

	for i := lenght; i > 0; i-- {
		newList = append(newList, il[i-1])
	}

	return newList
}

//Append ...
func (il IntList) Append(originalList IntList) IntList {
	for _, item := range originalList {
		il = append(il, item)
	}
	return il
}

//Concat ...
func (il IntList) Concat(list []IntList) IntList {
	for _, items := range list {
		for _, item := range items {
			il = append(il, item)
		}
	}
	return il
}

//Map ...
func (il IntList) Map(fn unaryFunc) IntList {
	newList := IntList{}
	for _, item := range il {
		newList = append(newList, fn(item))
	}
	return newList
}

//Filter ...
func (il IntList) Filter(fn predFunc) []int {
	newList := make(IntList, 0, len(il))
	for _, item := range il {
		if fn(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

//Foldr ...
func (il IntList) Foldr(fn binFunc, initial int) int {
	for i := len(il) - 1; i >= 0; i-- {
		initial = fn(il[i], initial)
	}
	return initial
}

//Foldl ...
func (il IntList) Foldl(fn binFunc, initial int) int {
	for _, item := range il {
		initial = fn(initial, item)
	}
	return initial
}
