package ArrangingCoins

import (
	"awesomeProject/Common"
	"testing"
)

func TestArrangingCoins(t *testing.T) {
	Common.AssertEqual(ArrangingCoins(3), 2, t)
	Common.AssertEqual(ArrangingCoins(5), 2, t)
	Common.AssertEqual(ArrangingCoins(8), 3, t)
}
