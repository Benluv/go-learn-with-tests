package structsmethodsinterfaces

import "testing"

func TestPermiter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(62.0, 6.0)
	want := 372.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
