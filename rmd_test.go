package slices

import (
	"reflect"
	"testing"
)

func TestRmDInt(t *testing.T) {
	testCases := []struct {
		slice  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5, 1, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 3, 4, 5, 1, 2}, []int{1, 3, 4, 5, 2}},
	}
	for _, test := range testCases {
		newSlice := RmDInt(test.slice)
		if ok := reflect.DeepEqual(newSlice, test.expect); !ok {
			t.Fatalf("RmDInt slices: %+v, want: %+v, but got: %+v",
				test.slice, test.expect, newSlice)
		}
	}
}

func TestRmDString(t *testing.T) {
	testCases := []struct {
		slice  []string
		expect []string
	}{
		{[]string{"1", "2", "3", "4", "5", "1", "2"}, []string{"1", "2", "3", "4", "5"}},
		{[]string{"1", "3", "4", "5", "1", "2"}, []string{"1", "3", "4", "5", "2"}},
	}
	for _, test := range testCases {
		newSlice := RmDString(test.slice)
		if ok := reflect.DeepEqual(newSlice, test.expect); !ok {
			t.Fatalf("RmDString slices: %+v, want: %+v, but got: %+v",
				test.slice, test.expect, newSlice)
		}
	}
}

func BenchmarkRmDInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmDInt([]int{1, 2, 3, 4, 5, 1, 2})
	}
}

func BenchmarkRmDString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RmDString([]string{"1", "2", "3", "4", "5", "1", "2"})
	}
}
