package intslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intSlice(ints ...int) []int {
	return ints
}

func pointerFunc(i int) func() *int {
	return func() *int {
		return &i
	}
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
			actual := Filter(test.ints, test.bfn)

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
			actual := Some(test.ints, test.bfn)

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
			actual := Every(test.ints, test.bfn)

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
			actual := Map(test.ints, test.ifn)

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
			actual := Find(test.ints, test.bfn)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		bfn      BoolFn
		expected func() *int
	}{
		{
			name: "Found returns index of int",
			ints: intSlice(101, 200, 150, 1000),
			bfn: func(i int, _ int, _ []int) bool {
				return i == 150
			},
			expected: func() *int {
				i := 2
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
			actual := FindIndex(test.ints, test.bfn)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		includes int
		expected bool
	}{
		{
			name:     "Found returns index of int",
			ints:     intSlice(101, 200, 150, 1000),
			includes: 200,
			expected: true,
		},
		{
			name:     "Not found returns nil",
			ints:     intSlice(1, 200, 150, 101),
			includes: 5,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Includes(test.ints, test.includes)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		includes int
		expected func() *int
	}{
		{
			name:     "Found returns index of int",
			ints:     intSlice(101, 200, 150, 1000),
			includes: 101,
			expected: func() *int {
				i := 0
				return &i
			},
		},
		{
			name:     "Not found returns nil",
			ints:     intSlice(1, 200, 150, 101),
			includes: 5,
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual :=IndexOf(test.ints, test.includes)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		join     string
		expected string
	}{
		{
			name:     "Found returns index of int",
			ints:     intSlice(101, 200, 150, 1000),
			join:     ",",
			expected: "101,200,150,1000",
		},
		{
			name:     "Not found returns nil",
			ints:     intSlice(1, 200, 150, 101),
			join:     "/",
			expected: "1/200/150/101",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Join(test.ints, test.join)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		expected func() *int
	}{
		{
			name: "Found returns index of int",
			ints: intSlice(101, 200, 150, 1000),
			expected: func() *int {
				i := 1000
				return &i
			},
		},
		{
			name:     "Not found returns nil",
			ints:     intSlice(),
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual :=Last(test.ints)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		expected func() *int
	}{
		{
			name: "Returns first int",
			ints: intSlice(101, 200, 150, 1000),
			expected: func() *int {
				i := 101
				return &i
			},
		},
		{
			name:     "Empty list returns nil",
			ints:     intSlice(),
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := First(test.ints)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestAt(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		at       int
		expected func() *int
	}{
		{
			name: "Returns *int at index",
			ints: intSlice(101, 200, 150, 1000),
			at:   2,
			expected: func() *int {
				i := 150
				return &i
			},
		},
		{
			name:     "Empty list returns nil",
			ints:     intSlice(),
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := At(test.ints, test.at)

			if actual == nil {
				assert.Nil(t, test.expected())
				return
			}

			assert.Equal(t, *test.expected(), *actual)
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name    string
		ints    []int
		prepend int
	}{
		{
			name:    "Prepends int",
			ints:    intSlice(101, 200, 150, 1000),
			prepend: 2,
		},
		{
			name:    "Empty list returns nil",
			ints:    intSlice(),
			prepend: 100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Prepend(test.ints, test.prepend)

			assert.Equal(t, *From(actual).First(), test.prepend)
		})
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		expected bool
	}{
		{
			name:     "Prepends int",
			ints:     intSlice(101, 200, 150, 1000),
			expected: false,
		},
		{
			name:     "Empty list returns nil",
			ints:     intSlice(),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Empty(test.ints)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		expected []int
	}{
		{
			name:     "Prepends int",
			ints:     intSlice(101, 200, 150, 1000, 5),
			expected: intSlice(5, 101, 150, 200, 1000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Sort(test.ints)

			for i := range actual {
				assert.Equal(t, test.expected[i], actual[i])
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
	}{
		{
			name:     "Prepends int",
			ints:     intSlice(101, 200, 150, 1000, 5),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := make([]int, 0, len(test.ints))
			actual := From(test.ints).Copy(expected)

			for i := range actual {
				assert.Equal(t, expected[i], actual[i])
			}
		})
	}
}
