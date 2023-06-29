package geo

type UTM struct {
}

func (u UTM) String() string {
	//TODO implement me
	panic("implement me")
}

func (u UTM) Decimal() float64 {
	//TODO implement me
	panic("implement me")
}

func (u UTM) Direction() Direction {
	//TODO implement me
	panic("implement me")
}

var _ Coordinate = UTM{}
