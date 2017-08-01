package a

import (
	"testing"
)

func TestAfunc(t *testing.T) {
	ret := Afunc("hoge")
	if ret != "A:hoge" {
		t.Errorf("error!")
	}
}
