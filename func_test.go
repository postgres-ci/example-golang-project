package test

import (
	"testing"
)

func TestSum(t *testing.T) {

	type values struct{ a, b int }

	assets := map[int]values{
		0: {0, 0},
		2: {1, 1},
		4: {2, 2},
	}

	for expected, v := range assets {

		if actual := Sum(v.a, v.b); expected != actual {

			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}
	}
}

func BenchmarkSum(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Sum(2, 2)
	}
}
