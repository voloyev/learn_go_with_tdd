package main

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 22.0}

	rectangle.Perimeter()
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want	%g want", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func BenchmarkArea(t *testing.B) {
	areaTests := []struct {
		shape Shape
	}{
		{Rectangle{12, 6}},
		{Circle{10}},
		{Triangle{12, 6}},
	}

	for _, tt := range areaTests {
		tt.shape.Area()
	}
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{12.0, 6.0}

	fmt.Printf("%.1f", rectangle.Area())
	// Output: 72.0
}

func ExampleRectangle_Perimeter() {
	rectangle := Rectangle{12.0, 6.0}
	fmt.Printf("%.1f", rectangle.Perimeter())
	// Output: 36.0
}

func ExampleCircle_Area() {
	circle := Circle{10.0}
	fmt.Printf("%.1f", circle.Area())
	// Output: 314.2
}

func ExampleTriangle_Area() {
	triangle := Triangle{10.0, 22.0}

	fmt.Printf("%.1f", triangle.Area())
	// Output: 110.0
}
