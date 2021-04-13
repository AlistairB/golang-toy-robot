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

func move(robot model.Robot, maxCoordinate model.Coordinate) model.Robot {
	// var newCoordinate model.Coordinate

	switch robot.Facing {
	case model.North:
		robot.Position.Y += 1
	// case model.West:
	// 	result = model.South
	// case model.South:
	// 	result = model.East
	// case model.East:
	// 	result = model.North
	}

	return robot
}

func coordinateInBounds(coordinate model.Coordinate, maxCoordinate model.Coordinate) bool {
	return coordinate.X >= 1 && coordinate.X <= maxCoordinate.X &&
		   coordinate.Y >= 1 && coordinate.Y <= maxCoordinate.Y
}