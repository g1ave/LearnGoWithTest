package rectangle

import "testing"

func TestPerimeter(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("want '%.2f' got '%.2f'", want, got)
		}
	}

	t.Run("caculate the perimeter of rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		assertCorrectMessage(t, got, want)
	})

}
