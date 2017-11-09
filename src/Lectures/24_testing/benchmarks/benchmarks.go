package benchmarks

func FacultyRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * FacultyRecursive(n-1)
}

func FacultyLoop(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
