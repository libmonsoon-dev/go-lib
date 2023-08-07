package geo

//go:generate stringer -type Direction -trimprefix Direction
type Direction byte

const (
	DirectionUndefined Direction = iota
	DirectionNorth
	DirectionSouth
	DirectionWest
	DirectionEast

	N = DirectionNorth
	S = DirectionSouth
	W = DirectionWest
	E = DirectionEast
)

func (d Direction) Rune() rune {
	if d < DirectionNorth || d > DirectionEast {
		return '\000'
	}

	return rune(_Direction_name[_Direction_index[d]])
}
