package conv

import "testing"

func BenchmarkParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wait := make(chan int)
		Parallel(10, wait)
		<-wait
	}
}

func BenchmarkParallel20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wait := make(chan int)
		Parallel(20, wait)
		<-wait
	}
}

func BenchmarkParallel30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wait := make(chan int)
		Parallel(30, wait)
		<-wait
	}
}

func BenchmarkParallel40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wait := make(chan int)
		Parallel(40, wait)
		<-wait
	}
}

func BenchmarkParallel50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wait := make(chan int)
		Parallel(50, wait)
		<-wait
	}
}

func BenchmarkLinear10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(10)
	}
}

func BenchmarkLinear20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(20)
	}
}

func BenchmarkLinear30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(30)
	}
}

func BenchmarkLinear40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(40)
	}
}

func BenchmarkLinear50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(50)
	}
}
