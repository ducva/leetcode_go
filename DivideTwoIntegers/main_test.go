package DivideTwoIntegers
import "testing"

func assertEqual(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Wrong: expected: %d, actual: %d", expected, actual)
	}
}
func TestSubtract(t *testing.T) {
	var ans = Subtract(2, 1)
	assertEqual(1, ans, t)
	assertEqual(100, Subtract(101, 1), t)
}
func TestAdd(t *testing.T) {
	var ans = Add(1,1)
	if ans != 2 {
		t.Errorf("Wrong: expected: %d, actual: %d", 2, ans)
	}
	ans = Add(1,0)
	if ans != 1 {
		t.Errorf("Wrong: expected: %d, actual: %d", 1, ans)
	}
	ans = Add(100,200)
	if ans != 300 {
		t.Errorf("Wrong: expected: %d, actual: %d", 300, ans)
	}
}

func TestMul(t *testing.T) {
	assertEqual(2, Mul(1, 2), t)
	assertEqual(0, Mul(0, 2), t)
	assertEqual(15, Mul(3, 5), t)
	assertEqual(200, Mul(25, 8), t)
	assertEqual(-15, Mul(3, -5), t)
}

func TestCase1(t *testing.T) {
	var ans = Divide(10, 3)

	if ans != 3 {
		t.Errorf("Wrong: expected: %d, actual: %d", 3, ans)
	}
	ans = Divide(7, -3)
	if ans != -2 {
		t.Errorf("Wrong: expected: %d, actual: %d", -2, ans)
	}
	ans = Divide(-1, 1)
	if ans != -1 {
		t.Errorf("Wrong: expected: %d, actual: %d", -1, ans)
	}

	ans = Divide(-2147483648, -1)
	if ans != 2147483647 {
		t.Errorf("Wrong: expected: %d, actual: %d", 2147483647, ans)
	}

	ans = Divide(2147483648, -1)
	if ans != -2147483648 {
		t.Errorf("Wrong: expected: %d, actual: %d", -2147483648, ans)
	}

}
