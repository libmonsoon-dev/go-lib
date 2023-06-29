package geo

import "strconv"

func NewDecimalLatitude(value float64) Decimal {
	return Decimal{Value: value}
}

func NewDecimalLongitude(value float64) Decimal {
	return Decimal{value, true}
}

type Decimal struct {
	Value       float64
	isLongitude bool
}

func (d Decimal) String() string {
	return strconv.FormatFloat(d.Value, 'g', -1, 64)
}

func (d Decimal) Decimal() float64 {
	return d.Value
}

func (d Decimal) Direction() Direction {
	switch {
	case !d.isLongitude && d.Value > 0:
		return DirectionNorth
	case !d.isLongitude:
		return DirectionSouth
	case d.isLongitude && d.Value > 0:
		return DirectionEast
	case d.isLongitude:
		return DirectionWest
	default:
		panic("uncharitable")
	}
}

var _ Coordinate = Decimal{}
