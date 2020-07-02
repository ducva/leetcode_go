package ReverseString

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	var input = []byte{'h', 'e', 'l', 'l', 'o'}
	var expected = []byte{'o', 'l', 'l', 'e', 'h'}
	ReverseString(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Failed. Actual: %v", input)
	}
}
