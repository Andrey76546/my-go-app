package main

import "testing"

func TestSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "opposite numbers", a: -1, b: 1, expected: 0},
		{name: "zero values", a: 0, b: 0, expected: 0},
		{name: "negative numbers", a: -5, b: -7, expected: -12},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if result := Sum(test.a, test.b); result != test.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}
