package goutil

import "testing"

func TestDMapAdd(t *testing.T) {
	d := new(DMap)
	d.Add("key", 5)
	Assert(t, d.Length() == 1, "wrong length")
	d.Add("key", 5)
	Assert(t, d.Length() == 1, "wrong length")
	d.Add("key2", 5)
	Assert(t, d.Length() == 2, "wrong length")
}

func TestDMapRemove(t *testing.T) {
	d := new(DMap)
	d.Add("key", 5)
	Assert(t, d.Length() == 1, "wrong length")
	d.Remove("key")
	Assert(t, d.Length() == 0, "wrong length")
	d.Add("key", 5)
	Assert(t, d.Length() == 1, "wrong length")
	d.Add("key2", 5)
	Assert(t, d.Length() == 2, "wrong length")
}

func TestDMapGet(t *testing.T) {
	d := new(DMap)
	d.Add("key", 5)
	Assert(t, d.Length() == 1, "wrong length")
	a := d.Get("key")
	Assert(t, a == 5, "wrong value")
}

func TestDMapNext(t *testing.T) {
	d := new(DMap)
	d.Add("key", 5)
	d.Add("key2", 6)
	Assert(t, d.Length() == 2, "wrong length")
	a := d.Next()
	Assert(t, a == 5, "wrong 1st value")
	a = d.Next()
	Assert(t, a == 6, "wrong 2nd value")
}
