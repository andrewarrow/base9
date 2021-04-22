package nonary

/*

  [ 8 ] [ 8 ] [ 8 ]
  [ 1 ] [ 0 ] [ 0 ] [ 0 ]

  increment yourself and if you spill over,
	tell digit to your left they are now + 1, u goto 0
	if there is no digit to your left,
	there is now, and it's a one

	&{[1 8]}
&{[2 1 0]}

you have to consume the 1, you don't assume this is case for whole new digit
	&{[2 0]}

*/

type Int struct {
	List []byte
}

func NewInt(a int) *Int {
	n := Int{}
	n.List = []byte{0}
	return &n
}
func (i *Int) Inc() *Int {

	// i have no idea how long the i.List is
	// but I can recurse
	// cases are
	// 1. just incr me
	// 2. set me to zero but call Inc again on digit to my left
	for j := len(i.List) - 1; j > -1; j-- {
		toMyLeft := i.fromDigitPerspective(j)
		if toMyLeft == false {
			i.List = append([]byte{1}, i.List...)
			break
		}
	}
	return i
}

func (i *Int) fromDigitPerspective(digit int) bool {
	val := i.List[digit]

	if val == 8 {
		i.List[digit] = 0
		return false
	}
	i.List[digit]++
	return true
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
