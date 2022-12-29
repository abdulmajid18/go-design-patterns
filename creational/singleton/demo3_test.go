package singleton

import (
	"testing"
)

func TestDemo3(t *testing.T) {
	s1 := New()
	s1["Abdul"] = "Majid"
	s2 := New()

	if s2["Abdul"] != "Majid" {
		t.Error("Expected Majid ")
	}
}
