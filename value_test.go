package slices

import "testing"

func TestIntValue(t *testing.T) {
	testCases := []struct {
		slice  []int
		index  int
		expect int
	}{
		{[]int{1, 2, 3}, -1, 0},
		{[]int{1, 2, 3}, 0, 1},
		{[]int{1, 2, 3}, 1, 2},
		{[]int{1, 2, 3}, 2, 3},
		{[]int{1, 2, 3}, 3, 0},
		{[]int{1, 2, 3}, 4, 0},
	}
	for _, test := range testCases {
		if intValue := IntValue(test.slice, test.index); intValue != test.expect {
			t.Fatalf("IntValue slices: %+v, index: %d, want: %d, but got: %d",
				test.slice, test.index, test.expect, intValue)
		}
	}
}

func TestStringValue(t *testing.T) {
	testCases := []struct {
		slice  []string
		index  int
		expect string
	}{
		{[]string{"1", "2", "3"}, -1, ""},
		{[]string{"1", "2", "3"}, 0, "1"},
		{[]string{"1", "2", "3"}, 1, "2"},
		{[]string{"1", "2", "3"}, 2, "3"},
		{[]string{"1", "2", "3"}, 3, ""},
		{[]string{"1", "2", "3"}, 4, ""},
	}
	for _, test := range testCases {
		if intValue := StringValue(test.slice, test.index); intValue != test.expect {
			t.Fatalf("StringValue slices: %+v, index: %d, want: %s, but got: %s",
				test.slice, test.index, test.expect, intValue)
		}
	}
}

func BenchmarkIntValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntValue([]int{1, 2, 3}, 2)
	}
}

func BenchmarkStringValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringValue([]string{"1", "2", "3"}, 2)
	}
}
