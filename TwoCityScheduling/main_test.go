package TwoCityScheduling

import (
	"awesomeProject/Common"
	"testing"
)

func TestTwoCityCost(t *testing.T) {
	var input = [][]int{{10,20},{30,200},{400,50},{30,20}}
	var expected = 110

	var actual = TwoCityCost(input)
	if actual != expected {
		t.Errorf("Failed: Actual: %d, expected: %d", actual, expected)
	}
	
	actual = TwoCityCost([][]int{{259,770},{448,54},{926,667},{184,139},{840,118},{577,469}})
	Common.AssertEqual(1859, actual, t)

}