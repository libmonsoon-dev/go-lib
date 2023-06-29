package geo

import "fmt"

type Coordinate interface {
	fmt.Stringer
	//TODO: fmt.Scanner

	Decimal() float64
	Direction() Direction
}
