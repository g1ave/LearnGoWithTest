package shape

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10, 10}, want: 100.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "tritangle", shape: Tritangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g but want %g", tt.shape, got, tt.want)
			}
		})
	}
}
