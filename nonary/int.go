package nonary

/*

  [ 8 ] [ 8 ] [ 8 ]
  [ 1 ] [ 0 ] [ 0 ] [ 0 ]

  increment yourself and if you spill over,
	tell digit to your left they are now + 1, u goto 0
	if there is no digit to your left,
	there is now, and it's a one

*/

type Int struct {
	List []byte
}

func NewInt(a int) *Int {
	n := Int{}
	n.List = []byte{0}
	return &n
}
func (i *Int) Add(a int) *Int {
	return i
}

/*
func (i *Int) Add2(a int) *Int {
	end := len(i.List) - 1
	i.List[end].Val++
	if i.List[end].Val > 8 {
		i.List[end].Val = 0
		i.HandleBump(len(i.List))
	}
	return i
}

func (i *Int) HandleBump(limit int) {
	if len(i.List) == limit && i.List[0].Val == 8 {
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
		end := len(i.List) - 1
		i.List[end].Val++
		if i.List[end].Val > 8 {
			i.HandleBump(len(i.List))
		}
	}
}*/
