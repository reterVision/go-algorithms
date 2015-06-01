package jumphashing

import (
	"math/rand"
	"testing"
)

func TestJumpHash(t *testing.T) {
	result := map[int]int{}
	buckets := rand.Int() % 100
	for i := 0; i < 10000; i++ {
		key := rand.Int() % 10
		bucket := JumpHash(key, buckets)
		if _, ok := result[key]; ok {
			if result[key] != bucket {
				t.Errorf("Same key %d, different bucket %d", key, bucket)
			}
		} else {
			result[key] = bucket
		}
	}
}

func BenchmarkJumpHash(b *testing.B) {
	buckets := rand.Int() % 100
	for i := 0; i < b.N; i++ {
		key := rand.Int() % 10
		_ = JumpHash(key, buckets)
	}
}
