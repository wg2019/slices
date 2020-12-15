package slices

import (
	"testing"
)

func TestRmDInt(t *testing.T) {
	newSlice := RmDInt([]int{1, 2, 3, 4, 5, 1, 2})
	t.Logf("newSlice: %+v", newSlice)
}

func TestRmDString(t *testing.T) {
	newSlice := RmDString([]string{"1", "2", "3", "4", "5", "1", "2"})
	t.Logf("newSlice: %+v", newSlice)
}
