package main

import "testing"
import "slices"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Error("got:", got, "expected:", want)
	}
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}

	expected := []int{3, 9}
	got := SumAll(slice1, slice2)

	if !slices.Equal(expected, got) {
		t.Error("expected:", expected, "got:", got)
	}
}

func TestSumTails(t *testing.T) {
	t.Run("test non empty slices", func(t *testing.T) {
		expected := []int{2, 9}
		got := SumAllTails([]int{1, 2}, []int{0, 9})

		if !slices.Equal(expected, got) {
			t.Error("expected:", expected, "got:", got)
		}
	})

	t.Run("test empty slice", func(t *testing.T) {
		expected := []int{0, 9}
		got := SumAllTails([]int{}, []int{3, 4, 5})

		if !slices.Equal(got, expected) {
			t.Error("expected:", expected, "got:", got)
		}
	})
}
