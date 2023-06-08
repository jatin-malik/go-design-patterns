package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 10
	expected := 15
	v := sum(a, b)
	if expected != v {
		t.Errorf("%d+%d should be %d,got %d", a, b, expected, v)
	}
}
