package main

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("b")
	fmt.Println(repeat)
	//Output: bbbbb
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
