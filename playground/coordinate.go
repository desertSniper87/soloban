package playground

import (
	"fmt"
	"reflect"
)
/*
* Main Coordinate Datatype
*/
type Coordiante struct {
	Row int
	Col int
}

func (c Coordiante) HashCode() uint  {
	row := uint(c.Row)
	col := uint(c.Col)
	return row * 1000 + col
}

func (coordiante1 Coordiante) equals(coordiante2 Coordiante) bool {
	return coordiante1.Col == coordiante2.Col && coordiante1.Row == coordiante2.Row
}


type SetCoord map[Coordiante]bool

func (s SetCoord) Add(coordiante Coordiante) {
	s[coordiante] = true
}

func (s SetCoord) Remove(coordiante Coordiante) {
	delete(s, coordiante)
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
	if set[coordiante] == true {
		return true
	}
	return false
}

func (set SetCoord) Print() {
	for k, _ := range set {
		fmt.Println(k)
	}
}

type StackCoord []Coordiante

func (s StackCoord) Push(coord Coordiante) StackCoord {
	return append(s, coord)
}

func (s StackCoord) Pop() (Coordiante, StackCoord) {
	n := len(s) - 1
	return s[n], s[:n]
}
