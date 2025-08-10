package main

import (
	"testing"
)

func BenchmarkAccessArrayByRow(b *testing.B) {
	var myArr [3][5]int
	b.ReportAllocs()
	b.ResetTimer()
	for k := 0; k < b.N; k++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 5; j++ {
				myArr[i][j] = i*i + j*j
			}
		}
	}
}

func BenchmarkAccessArrayByCol(b *testing.B) {
	var myArr [3][5]int
	b.ReportAllocs()
	b.ResetTimer()
	for k := 0; k < b.N; k++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				myArr[j][i] = i*i + j*j
			}
		}
	}
}
