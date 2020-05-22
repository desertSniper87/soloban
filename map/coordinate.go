package _map

type coordiante struct {
	row int
	col int
}



func (coordiante1 coordiante) equals(coordiante2 coordiante) bool {
	return coordiante1.col == coordiante2.col && coordiante1.row == coordiante2.row
}
