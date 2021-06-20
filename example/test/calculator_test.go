package main

import "testing"

func TestCalcSuccess(t *testing.T) {
	expected := 3
	actual := Calculate(1, 2)

	if expected != actual {
		t.Error("not matched!")
	}
}

func TestCalcFail(t *testing.T) {
	expected := 2
	actual := Calculate(1, 1)

	if expected != actual {
		t.Error("not matched!")
	}
}
