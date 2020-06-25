package Common

import "testing"

func AssertEqual(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Errorf("Wrong: expected: %s, actual: %s", expected, actual)
	}
}