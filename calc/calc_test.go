package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	ret := Sum(1, 2)
	if ret != 3 {
		t.Errorf("error!")
	}
}
