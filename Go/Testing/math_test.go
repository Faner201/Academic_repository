package math

import (
	"testing"
)

func TestABC(t *testing.T) {
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1)
	}
	if Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0)
	}
	if Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1)
	}
}

func TestAbcSub(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if Abs(1) < 0 {
			t.Error("Negative value found in abs()")
		}
	}) 
}

func TestSkip(t *testing.T) {
	getenv := os.Getenv("Gopath")
	if len(getenv) == 0 {
		t.Skip("Skipping test because Gopath isn't set")
	}
	t.Log("Test with Gopath", getenv)
}

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("Cleanup")
	})
	t.Log("Running some test")
}

//Method that parallels your tests
t.Paralel()