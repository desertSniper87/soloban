package playground

import (
	"reflect"
)
/*
* Main Coordinate Datatype
*/
type Coordiante struct {
	Row int
	Col int
}

func (c Coordiante) HashCode() int  {
	return c.Row * 1000 + c.Row
}

func (coordiante1 Coordiante) equals(coordiante2 Coordiante) bool {
	return coordiante1.Col == coordiante2.Col && coordiante1.Row == coordiante2.Row
}


type SetCoord map[Coordiante]bool

func (s SetCoord) Add(coordiante Coordiante) {
	s[coordiante] = true
}

func (set1 SetCoord) Equal(set2 SetCoord) bool {
	eq := reflect.DeepEqual(set1, set2)
	if eq {
		return true
	} else {
		return false
	}
}

func (set SetCoord) Contains(coordiante Coordiante) bool {
	//for s, _ := range set{
	//	fmt.Println(s)
	//}

	if set[coordiante] == true {
		return true
	}
	return false
}

type StackCoord []Coordiante

func (s StackCoord) Push(coord Coordiante) StackCoord {
	return append(s, coord)
}

func (s StackCoord) Pop() (Coordiante, StackCoord) {
	n := len(s) - 1
	return s[n], s[:n]
}
