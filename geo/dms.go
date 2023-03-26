package geo

type DMS struct {
	Degrees, Minutes, Seconds int
	DirectionFiled            Direction
}

func (c DMS) Direction() Direction {
	return c.DirectionFiled
}

func (c DMS) Decimal() float64 {
	// TODO: multiply to direction sign
	// TODO: use decimal for calculation
	return float64(c.Degrees) + (float64(c.Minutes) / 60) + (float64(c.Seconds) / 3600)
}

var _ Coordinate = DMS{}
