package integers

import (
	"testing"
	"fmt" 
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Error("got:", sum, "expected:", expected)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Error("got:", got, "expected:", want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output: 6
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// output: aaaaa
}
