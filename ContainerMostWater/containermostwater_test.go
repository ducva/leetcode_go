package containermostwater

import "testing"

func TestCase1(t *testing.T) {
	ans := MaxArea([]int{2,3,4,5,18,17,6})
	if ans != 17 {
		t.Errorf("[Wrong] expected: %d, actual: %d.", 17, ans)
	}

}