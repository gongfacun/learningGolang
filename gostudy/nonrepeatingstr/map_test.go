package nonrepeatingstr

import "testing"

func TestMap(t *testing.T) {
	tests := []struct {
		s     string
		count int
	}{
		{"adada", 2},
		{"dfsfs", 3},
		{"adadadbc", 4},
		{"", 0},
		{"aaaaa", 1},
		{"一二三二一", 3},
		{"需要传入参数的情况，当然不需要传入参数也可以用这种方法", 21},
	}

	for _, test := range tests {

		if actual := searchMaxlong(test.s); actual != test.count {

			t.Errorf("result %d of input %s  expected: %d", actual, test.s, test.count)

		}

	}

}

func BenchmarkMap(t *testing.B) {

	s := "需要传入参数的情况，当然不需要传入参数也可以用这种方法"
	count := 21

	for i := 0; i < t.N; i++ {

		if actual := searchMaxlong(s); actual != count {

			t.Errorf("result %d of input %s  expected: %d", actual, s, count)

		}
	}

}
