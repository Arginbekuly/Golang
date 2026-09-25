package main

import "testing"

var data = []byte("  \r\n  Hello, World!  \r\n  ")

func BenchmarkProcessPayloadOriginal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = processPayloadOriginal(data)

	}
}

func BenchmarkProcessPayload(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = processPayload(data)
	}
}
