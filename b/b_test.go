package b

import (
	"testing"
)

func TestBfunc(t *testing.T) {
	ret := Bfunc("hoge")
	if ret != "B:hoge" {
		t.Errorf("error!")
	}
}
