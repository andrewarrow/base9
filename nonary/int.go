package nonary

type Int struct {
	List []byte
}

func NewInt(a int) *Int {
	n := Int{}
	n.List = []byte{0}
	return &n
}

func (i *Int) Add(a int) {
	if a > 0 {
		for j := 0; j < a; j++ {
			i.Incr()
		}
		return
	}
	for j := 0; j < a*-1; j++ {
		i.Decr()
	}
}

func (i *Int) Incr() {
	overflow := false
	for j := len(i.List) - 1; j > -1; j-- {
		i.List[j]++

		if i.List[j] == 9 {
			i.List[j] = 0
			overflow = true
		} else {
			overflow = false
			break
		}
	}
	if overflow {
		i.List = append([]byte{1}, i.List...)
	}
}
func (i *Int) Decr() {
	underflow := false
	for j := len(i.List) - 1; j > -1; j-- {

		if i.List[j] == 0 {
			i.List[j] = 8
			underflow = true
		} else {
			i.List[j]--
			break
		}
	}
	if underflow && i.List[0] == 0 {
		newList := []byte{}
		for j, thing := range i.List {
			if j > 0 {
				newList = append(newList, thing)
			}
		}
		i.List = newList
	}
}
