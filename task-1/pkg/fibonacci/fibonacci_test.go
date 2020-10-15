package fibonacci

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := map[int]int{
		1:  0,
		2:  1,
		3:  1,
		5:  3,
		10: 34,
	}

	for k, v := range testCases {
		got := Calculate(k)

		if got != v {
			t.Errorf("Error got: %v, want:%v", got, v)
		}
	}
}
