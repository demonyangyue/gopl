package ex1_3

import "testing"

func BenchmarkJoin(b *testing.B) {
	data := []string{"hi", "yue"}

	for i := 0 ; i < b.N ; i++ {
		join(data)
	}
}

