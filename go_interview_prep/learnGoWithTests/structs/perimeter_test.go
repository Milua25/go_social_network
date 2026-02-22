package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10.0, Height: 20.0}
	got := rect.Perimeter()
	want := 60.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Area of rectangle", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 20.0}
		want := 200.0
		checkArea(t, rect, want)
	})

	t.Run("Area of circle", func(t *testing.T) {
		circle := Circle{radius: 5.0}
		want := 78.53981633974483
		checkArea(t, circle, want)
	})

	t.Run("Area of Shape", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{Width: 30, Height: 30}, 900.0},
			{Circle{radius: 10}, 314.1592653589793},
		}

		for _, tt := range areaTests {
			checkArea(t, tt.shape, tt.want)
		}
	})

	t.Run("Area of triangle", func(t *testing.T) {
		triangle := Triangle{Base: 10, Height: 20}
		want := 100.0
		checkArea(t, triangle, want)
	})
}
