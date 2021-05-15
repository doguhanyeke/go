package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{8.0, 12.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("Area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("Area of circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}