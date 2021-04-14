package main

import (
	"fmt"
	"golang-toy-robot/model"
	"golang-toy-robot/processing"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	table := model.Table{ MaxCoordinate: model.Coordinate{5,5}, Robot: nil}

	fmt.Println("~~ Toy Robot ~~")

	for {
		readCommandAsText := ""
		fmt.Scanln(readCommandAsText)

		toRunCommand := processing.ParseCommand(readCommandAsText)

		table = processing.ProcessCommand(table, toRunCommand)
	}
	// filename := os.Getenv("ROBOT_FILE")
	// ui.Play(filename)
}
