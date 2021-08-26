package intslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func intSlice(ints ...int) []int {
	return ints
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		ints        []int
		bfn         BoolFn
		expectedLen int
	}{
		{
			name: "Filters out ints < 10",
			ints: intSlice(0, 2, 3, 40, 10, 11),
			bfn: func(i int, _ int, _ []int) bool {
				return i > 10
			},
			expectedLen: 2,
		},
		{
			name: "Filters out ints > 10",
			ints: intSlice(0, 2, 3, 40, 10, 11),
			bfn: func(i int, _ int, _ []int) bool {
				return i < 10
			},
			expectedLen: 3,
		},
		{
			name: "Filters out ints != 10",
			ints: intSlice(10, 2, 3, 40, 10, 11, 10),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 10
			},
			expectedLen: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := From(test.ints).Filter(test.bfn)

			assert.Equal(t, test.expectedLen, len(actual))
		})
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		bfn      BoolFn
		expected bool
	}{
		{
			name: "Slice contains int == 100",
			ints: intSlice(1, 200, 150, 100),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 100
			},
			expected: true,
		},
		{
			name: "Slice does not contains int == 100",
			ints: intSlice(1, 200, 150, 101),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 100
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := From(test.ints).Some(test.bfn)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		bfn      BoolFn
		expected bool
	}{
		{
			name: "Every int > 100",
			ints: intSlice(101, 200, 150, 1000),
			bfn: func(i int, _ int, _ []int) bool {
				return i > 100
			},
			expected: true,
		},
		{
			name: "Every int not > 100",
			ints: intSlice(1, 200, 150, 101),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 100
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := From(test.ints).Every(test.bfn)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		ifn      IntFn
		expected []int
	}{
		{
			name: "i * 2",
			ints: intSlice(101, 200, 150, 1000),
			ifn: func(i int, _ int, _ []int) int {
				return i * 2
			},
			expected: intSlice(202, 400, 300, 2000),
		},
		{
			name: "i / 2",
			ints: intSlice(1, 200, 150, 101),
			ifn: func(i int, _ int, _ []int) int {
				return i / 2
			},
			expected: intSlice(0, 100, 75, 50),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := From(test.ints).Map(test.ifn)

			for i := range actual {
				assert.Equal(t, test.expected[i], actual[i])
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		bfn      BoolFn
		expected func() *int
	}{
		{
			name: "Found returns int",
			ints: intSlice(101, 200, 150, 1000),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 150
			},
			expected: func() *int {
				i := 150
				return &i
			},
		},
		{
			name: "Not found returns nil",
			ints: intSlice(1, 200, 150, 101),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 5
			},
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := From(test.ints).Find(test.bfn)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestFindIndex(t *testing.T) {

}

func TestIncludes(t *testing.T) {

}

func TestIndexOf(t *testing.T) {

}

func TestJoin(t *testing.T) {

}

func TestLast(t *testing.T) {

}

func TestFirst(t *testing.T) {

}

func TestAt(t *testing.T) {

}

func TestPrepend(t *testing.T) {

}

func TestEmpty(t *testing.T) {

}

func TestSort(t *testing.T) {

}

func TestCopy(t *testing.T) {

}
