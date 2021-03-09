package model

type Table struct {
	maxCoordinate Coordinate
	robot *Robot
}

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

type Robot struct {
	position Coordinate
	facing Direction
}

type Coordinate struct {
	x int
	y int
}
