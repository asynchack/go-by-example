package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("want -2, got %d\n", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d-%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("want %d, got %d", tt.want, ans)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
