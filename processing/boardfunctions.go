package processing

import (
	"golang-toy-robot/model"
	"fmt"
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
	newCoordinate := robot.Position

	switch robot.Facing {
	case model.North:
		newCoordinate.Y += 1
	case model.West:
		newCoordinate.X += 1
	case model.South:
		newCoordinate.Y -= 1
	case model.East:
		newCoordinate.X -= 1
	}

	if coordinateInBounds(newCoordinate, maxCoordinate) {
		robot.Position = newCoordinate
	}

	return robot
}

func coordinateInBounds(coordinate model.Coordinate, maxCoordinate model.Coordinate) bool {
	return coordinate.X >= 1 && coordinate.X <= maxCoordinate.X &&
		   coordinate.Y >= 1 && coordinate.Y <= maxCoordinate.Y
}

func reportRobot(robot model.Robot) string {
	var facingAsString string

	switch robot.Facing {
	case model.North:
		facingAsString = "North"
	case model.West:
		facingAsString = "West"
	case model.South:
		facingAsString = "South"
	case model.East:
		facingAsString = "East"
	}

	return fmt.Sprintf("%d,%d,%s", robot.Position.X, robot.Position.Y, facingAsString)
}