package model

type Table struct {
	maxCoordinate Coordinate
	robot         *Robot
}

func PlaceRobot(table Table, robot Robot) Table {
	return Table{table.maxCoordinate, &robot}
}
