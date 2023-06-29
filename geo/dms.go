package geo

import (
	"fmt"
)

type DMS struct {
	Degrees, Minutes, Seconds int
	DirectionFiled            Direction
}

func (c DMS) String() string {
	return fmt.Sprintf("%d° %d′ %d″ %c", c.Degrees, c.Minutes, c.Seconds, c.DirectionFiled.Rune())
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
