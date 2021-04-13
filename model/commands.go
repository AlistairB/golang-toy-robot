package model

type CommandType int

type Command interface {
    isCommand()
}

type PlaceCommand struct {
	Coordinate Coordinate
	Facing Direction
}

type MoveCommand struct {}
type LeftCommand struct {}
type RightCommand struct {}
type ReportCommand struct {}

func (pc PlaceCommand) isCommand() {}
func (pc MoveCommand) isCommand() {}
func (pc LeftCommand) isCommand() {}
func (pc RightCommand) isCommand() {}
func (pc ReportCommand) isCommand() {}