package main

import "testing"

func TestAddSuccess(t *testing.T) {
	expected := 3
	actual := Calculate(1, 2, "+")

	if expected != actual {
		t.Error("not matched!")
	}
}
