package IntegerToRoman

import (
	"awesomeProject/Common"
	"testing"
)

func TestInt2Roman(t *testing.T) {
	Common.AssertEqual("I", Int2Roman(1),  t)
	Common.AssertEqual("VI", Int2Roman(6),  t)
	Common.AssertEqual("IV", Int2Roman(4),  t)
	Common.AssertEqual("VII", Int2Roman(7),  t)
	Common.AssertEqual("IX", Int2Roman(9),  t)
}