package main

import (
	"bytes"
	"strings"
	"testing"
)

// 使用拼接符拼接字符串
func BenchmarkJoinStringUsePlus(b *testing.B) {
	strSlices := []string{"h", "e", "l", "l", "o"}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			var all string
			for _, str := range strSlices {
				all += str
			}
			_ = all
		}
	}
}

// 复用bytes.Buffer结构
func BenchmarkJoinStringUseBytesBufWithReuse(b *testing.B) {
	strSlices := []string{"h", "e", "l", "l", "o"}
	var bf bytes.Buffer
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			var all string
			for _, str := range strSlices {
				bf.WriteString(str)
			}
			all = bf.String()
			_ = all
			bf.Reset()
		}
	}
}

// 使用bytes.Buffer，未进行复用
func BenchmarkJoinStringUseBytesBufWithoutReuse(b *testing.B) {
	strSlices := []string{"h", "e", "l", "l", "o"}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			var all string
			var bf bytes.Buffer
			for _, str := range strSlices {
				bf.WriteString(str)
			}
			all = bf.String()
			_ = all
			//bf.Reset()
		}
	}
}

// 使用strings.Builder
func BenchmarkJoinStringUseStringBuilder(b *testing.B) {
	strSlices := []string{"h", "e", "l", "l", "o"}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			all := ""
			var strb strings.Builder
			for _, str := range strSlices {
				strb.WriteString(str)
			}
			all = strb.String()
			_ = all
			strb.Reset()
		}
	}
}
