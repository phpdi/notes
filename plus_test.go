package main

import (
	"testing"
)

func BenchmarkTestPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test1()
	}
}

func BenchmarkTestBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test2()
	}
}

func BenchmarkTestFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test3()
	}
}

func BenchmarkTestAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Test4()
	}
}
