package main
import (
	"testing"
)
func BenchmarkEscape(b *testing.B){
	for i:=0;i<b.N;i++ {
		escape(0.35+0.21i)
	}
}
