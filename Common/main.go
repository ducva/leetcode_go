package Common

import "testing"

func AssertEqual(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Wrong: expected: %s, actual: %s", expected, actual)
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
