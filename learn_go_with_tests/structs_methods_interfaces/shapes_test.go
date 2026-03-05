package main

import "testing"

func TestPerimeter(t *testing.T) {
	shape := Rectangle{10.0, 10.0}

	expected := 40.0
	got := Perimeter(shape)

	if got != expected {
		t.Error("expected:", expected, "got:", got)
	}
}

func TestArea(t *testing.T) {
	t.Run("testing with a rectangle", func(t *testing.T) {
		shape := Rectangle{30.0, 40.0}

		expected := 1200.0
		got := shape.Area()

		if expected != got {
			t.Error("expected:", expected, "got:", got)
		}
	})

	t.Run("testing with a circle", func(t *testing.T) {
		shape := Circle{10.0}

		expected := 314.1592653589793
		got := shape.Area()

		if got != expected {
			t.Error("expected:", expected, "got:", got)
		}
	})

	t.Run("using table driven tests", func(t *testing.T) {
		area_tests := []struct {
			shape Shape
			want float64
		} {
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{shape: Circle{Radius: 10}, want: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range area_tests {
			got := tt.shape.Area()

			if got != tt.want {
				t.Error("expected:", tt.want, "got:", got)	
			}
		}
	})
}
