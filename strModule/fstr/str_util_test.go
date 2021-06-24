package fstr

import "testing"

func TestTrim(t *testing.T) {
	expected := "hello world"
	if got := Trim(" hello world  "); got != expected {
		t.Errorf("Trim(\" hello world  \")=\"%s\"; want\"%s\"", got, expected)
	}

}
