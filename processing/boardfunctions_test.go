package processing

import (
	"golang-toy-robot/model"
	"testing"
)

func TestReportRobot(t *testing.T) {
	input := model.Robot{Position: model.Coordinate{1, 1}, Facing: model.North}
	result := reportRobot(input)
	expected := "1,1,NORTH"

	if result != expected {
		t.Errorf("Expected %s Got %s", expected, result)
	}
}