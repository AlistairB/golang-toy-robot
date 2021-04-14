package processing

import (
	"golang-toy-robot/model"
)

func ParseCommand(input string) model.Command {
	var result model.Command

	switch input {
	case "MOVE":
		result = model.MoveCommand{}
	case "LEFT":
		result = model.LeftCommand{}
	case "RIGHT":
		result = model.RightCommand{}
	case "Report":
		result = model.ReportCommand{}
	}

	return result
}