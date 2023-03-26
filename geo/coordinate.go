package geo

type Coordinate interface {
	Decimal() float64
	Direction() Direction
}
