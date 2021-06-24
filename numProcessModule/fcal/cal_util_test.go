package fcal

import "testing"

func TestCalInterest(t *testing.T) {
	expected := 10.0
	if got := CalInterest(100.0); got != expected {
		t.Errorf("CalInterest(100.0)=.2%f; want %.2f", got, expected)
	}
}
