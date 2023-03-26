package geo

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
