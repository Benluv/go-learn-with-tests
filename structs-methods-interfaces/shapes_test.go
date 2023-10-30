package structsmethodsinterfaces

import "testing"

func TestPermiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{62.0, 6.0}
	got := Area(rectangle)
	want := 372.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
