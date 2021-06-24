package fmath

import "testing"

func TestAddFloat(t *testing.T) {
	expected := 10.0
	if got := AddFloat(5.0, 5.0); got != expected {
		t.Errorf("AddFloat(5.0, 5.0)=%.2f; want %.2f", got, expected)
	}

}
