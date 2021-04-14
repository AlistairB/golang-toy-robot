package main

import (
	"fmt"
	"golang-toy-robot/model"
	"golang-toy-robot/processing"
)

func main() {
	table := model.Table{ MaxCoordinate: model.Coordinate{5,5}, Robot: nil}

	fmt.Println("~~ Toy Robot ~~")

	for {
		readCommandAsText := ""
		fmt.Scanln(readCommandAsText)

		toRunCommand := processing.ParseCommand(readCommandAsText)

		if toRunCommand != nil {
			table = processing.ProcessCommand(table, toRunCommand)
		}
	}
}
