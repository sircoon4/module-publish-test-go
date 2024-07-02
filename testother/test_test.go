package testother

import "testing"

func TestTest(t *testing.T) {
	want := "module testother"
	if got := TestOther(); got != want {
		t.Errorf("Test() = %q, want %q", got, want)
	}
}
