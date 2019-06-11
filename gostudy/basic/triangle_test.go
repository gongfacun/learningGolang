package basic

import "testing"

func TestTriangle(t *testing.T) {

	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{6, 8, 10},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}
	for _, test := range tests {

		if actual := colcTriangle(test.a, test.b); actual != test.c {

			t.Errorf("colctriangle(%d,%d); got %d expected %d", test.a, test.b, actual, test.c)
		}
	}
}
