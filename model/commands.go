package model

type CommandType int

const (
	Place CommandType = iota
	Move
	Left
	Right
	Report
)

type Command interface {
    getCommandType() CommandType
}

type PlaceCommand struct {
	coordinate Coordinate
}

type MoveCommand struct {}
type LeftCommand struct {}
type RightCommand struct {}
type ReportCommand struct {}

func (pc PlaceCommand) getCommandType() CommandType {
	return Place
}