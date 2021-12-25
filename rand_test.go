package go_example_module

import "testing"

func TestRand(t *testing.T) {
	val := MyRandInt64()
	if val == 0 {
		t.Error("value us zero")
	}
	if val < 0 {
		t.Errorf("Value is negative")
	}
}
