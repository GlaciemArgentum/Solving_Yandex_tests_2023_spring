package main

import (
	"testing"
)

// Проверить если перебором
func BenchmarkRealMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RealMain()
	}
}
