package goutil

import "testing"

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
