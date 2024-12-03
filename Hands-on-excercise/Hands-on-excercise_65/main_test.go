package main

import "testing"

func TestOperations(t *testing.T) {
	got_add := add(2, 3)
	want_add := 5
	if got_add != want_add {
		t.Errorf("got %d want %d", got_add, want_add)
	}
	got_sub := subtract(5, 3)
	want_sub := 2
	if got_sub != want_sub {
		t.Errorf("got %d want %d", got_sub, want_sub)
	}
	got_mul := multiply(2, 3)
	want_mul := 6
	if got_mul != want_mul {
		t.Errorf("got %d want %d", got_mul, want_mul)
	}
	got_div := divide(6, 3)
	want_div := 2.0
	if got_div != want_div {
		t.Errorf("got %f want %f", got_div, want_div)
	}
}
