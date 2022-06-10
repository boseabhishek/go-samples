package main

import "testing"

const creations = 20_000_000

func TestCopyIt(t *testing.T) {
	for i := 0; i < 20_000_000; i++ {
		_ = CreateCopy()
	}
}

func TestPointerIt(t *testing.T) {
	for i := 0; i < creations; i++ {
		_ = CreatePointer()
	}
}

func BenchmarkStackIt_Invoke_Main1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main1()
	}
}

func BenchmarkStackIt_Invoke_Main2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main2()
	}
}

func BenchmarkStackIt_Invoke_Main3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main3()
	}
}

func Benchmark_Invoke_CopyIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyIt()
	}
}

func Benchmark_Invoke_PointerIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointerIt()
	}
}
