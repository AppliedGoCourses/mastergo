package benchmarks

import "testing"

func BenchmarkFacultyRecursive(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"5!", 5},
		{"10!", 10},
		{"20!", 20},
		{"50!", 50},
		{"100!", 100},
		{"500!", 500},
		{"1000!", 1000},
	}
	for _, tt := range benchmarks {
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				FacultyRecursive(tt.num)
			}
		})
	}
}

func BenchmarkFacultyLoop(b *testing.B) {
	benchmarks := []struct {
		name string
		num  int
	}{
		{"5!", 5},
		{"10!", 10},
		{"20!", 20},
		{"50!", 50},
		{"100!", 100},
		{"500!", 500},
		{"1000!", 1000},
	}
	for _, tt := range benchmarks {
		b.Run(tt.name, func(b *testing.B) {
			for b.Loop() {
				FacultyLoop(tt.num)
			}
		})
	}
}
