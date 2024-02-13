package structs

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{25.0, 3.0}

// 	got := Area(rectangle)
// 	want := 75.0

// 	if got != want {
// 		t.Errorf("got %.2f, want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := Area(rectangle)
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := Area(circle)
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// }

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})

// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name      string
		shape     Shape
		shapeArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, shapeArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, shapeArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, shapeArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.shapeArea {
				t.Errorf("%#v got %g, though shape area is %g", tt.shape, got, tt.shapeArea)
			}
		})
	}
}
