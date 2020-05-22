package playground

type Coordiante struct {
	Row int
	Col int
}


func (coordiante1 Coordiante) equals(coordiante2 Coordiante) bool {
	return coordiante1.Col == coordiante2.Col && coordiante1.Row == coordiante2.Row
}
