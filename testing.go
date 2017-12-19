package goutil

import (
	"fmt"
	"testing"
)

// Assert wrapper for test failure
func Assert(t *testing.T, condition bool, err string) {
	if !condition {
		t.Log(err)
		t.Fail()
	}
}

// AssertNow wrapper for test failure
func AssertNow(t *testing.T, condition bool, err string) {
	if !condition {
		t.Log(err)
		t.FailNow()
	}
}

func AssertLength(t *testing.T, result int, expected int) {
	if result != expected {
		t.Log(fmt.Sprintf("wrong length: %d, expected: %d", result, expected))
	}
}
