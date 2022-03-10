package lock

import "testing"

func Test_New(t *testing.T) {
	l, err := FromFile("../yarn.lock")
	if err != nil {
		t.Fatal(err)
	}
	if l == nil {
		t.Fatalf("parse lock failed")
	}
}
