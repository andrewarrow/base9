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
	end := len(i.List) - 1
	i.List[end].Val = 0
	// 8, 0
	if len(i.List) == limit {
		fb := FancyBit{}
		fb.Val = 1
		i.List = append([]FancyBit{fb}, i.List...)
	} else {
		i.List[0].Val++
		if i.List[0].Val > 8 {
			i.List[0].Val = 0
			i.HandleBump(limit + 1)
		}
	}
}
