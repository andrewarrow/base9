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
	i.List[0].Val++
	if i.List[0].Val > 8 {
		i.List[0].Val = 0
		if len(i.List) == 1 {
			fb := FancyBit{}
			fb.Val = 1
			i.List = append([]FancyBit{fb}, i.List...)
		} else {
			i.List[1].Val++
		}
	}
	return i
}
