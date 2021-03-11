package processing

import (
	"golang-toy-robot/model"
)

func turnRight(direction model.Direction) model.Direction {
	var result model.Direction

	switch direction {
	case model.North:
		result = model.East
	case model.East:
		result = model.South
	case model.South:
		result = model.West
	case model.West:
		result = model.North
	}

	return result
}

func turnLeft(direction model.Direction) model.Direction {
	var result model.Direction

	switch direction {
	case model.North:
		result = model.West
	case model.West:
		result = model.South
	case model.South:
		result = model.East
	case model.East:
		result = model.North
	}

	return result
}
