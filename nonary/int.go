package nonary

type FancyBit struct {
	Val byte
}

type Int struct {
	List []FancyBit
}

func NewInt(a int) *Int {
	n := Int{}
	fb := FancyBit{}
	fb.Val = 0
	n.List = []FancyBit{fb}
	return &n
}

// 1 0 0 0 1
// 1 0 0 1 0
// 1 0 1 0 0
// 1 0 1 1 1
// 1 1 0 0 0
func (i *Int) Add(a int) *Int {
	end := len(i.List) - 1
	i.List[end].Val++
	// 8 8, 8, 9
	if i.List[end].Val > 8 {
		i.HandleBump(1)
	}
	return i
}

func (i *Int) HandleBump(limit int) {
	// 8, 9 - 1
	// 8, 0
	if len(i.List) == limit {
		fb := FancyBit{}
		fb.Val = 1
		foo := []FancyBit{}
		for j := 0; j < len(i.List); j++ {
			fb := FancyBit{}
			fb.Val = 0
			foo = append(foo, fb)
		}
		i.List = append([]FancyBit{fb}, foo...)
	} else {
		i.List[0].Val++
		if i.List[0].Val > 8 {
			i.HandleBump(limit + 1)
		}
	}
}
