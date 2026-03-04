package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Error("got:", got, "expected:", want)
	}
}
