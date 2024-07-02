package test

import "testing"

func TestTest(t *testing.T) {
	want := "module test"
	if got := Test(); got != want {
		t.Errorf("Test() = %q, want %q", got, want)
	}
}
