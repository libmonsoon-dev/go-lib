package geo

import (
	"fmt"

	"github.com/paulmach/orb"
)

type Point struct{ Latitude, Longitude Coordinate }

func (p Point) String() string {
	return fmt.Sprintf("%s, %s", p.Latitude, p.Longitude)
}

func (p Point) Point() orb.Point {
	p.check()

	return orb.Point{p.Longitude.Decimal(), p.Latitude.Decimal()}
}

func (p Point) check() {
	if p.Latitude.Direction() != N && p.Latitude.Direction() != S {
		panic(fmt.Sprintf("invalid latitude direction %v", p.Latitude.Direction()))
	}

	if p.Longitude.Direction() != E && p.Longitude.Direction() != W {
		panic(fmt.Sprintf("invalid longitude direction %v", p.Longitude.Direction()))
	}
}
