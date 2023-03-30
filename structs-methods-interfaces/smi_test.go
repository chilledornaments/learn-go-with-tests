package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		rectangle Rectangle
		want      float64
	}{
		{
			name:      "Square",
			rectangle: Rectangle{Width: 10.0, Height: 10.0},
			want:      40.0,
		},
		{
			name:      "Rectangle",
			rectangle: Rectangle{Width: 6.0, Height: 10.0},
			want:      32.0,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				got := tc.rectangle.Perimeter()

				if got != tc.want {
					t.Errorf("got '%f' but wanted '%f'", got, tc.want)
				}
			})
	}
}

func TestArea(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		want  float64
		shape Shape
	}{
		{
			name:  "Square",
			shape: Rectangle{Width: 10.0, Height: 10.0},
			want:  100.0,
		},
		{
			name:  "Rectangle",
			shape: Rectangle{Width: 5.0, Height: 10.0},
			want:  50.0,
		},
		{
			name:  "Circle",
			shape: Circle{Radius: 10},
			want:  314.1592653589793,
		},
		{
			name:  "Triangle",
			shape: Triangle{Base: 10, Height: 2.5},
			want:  12.5,
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				got := tc.shape.Area()

				if got != tc.want {
					t.Errorf("%#v got %g want %g", tc.shape, got, tc.want)
				}
			},
		)
	}
}
