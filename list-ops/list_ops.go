package listops

func main() {
	var x IntList
	x.Length()
}

//IntList ...
type IntList []int

//Length ...
func (il IntList) Length() int {
	return len(il)
}
