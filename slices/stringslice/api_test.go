package stringslice

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func strSlice(strs ...string) []string {
	return strs
}

func boolFnComposer(expected bool) BoolFn {
	return func(string, int, []string) bool {
		return expected
	}
}

func stringFnAccumulator(start string) StringFn {
	return func(i string, _ int, _ []string) string {
		return start + i
	}
}

func stringRef(s string) func() *string {
	return func() *string {
		return &s
	}
}

func intRef(i int) func() *int {
	return func() *int {
		return &i
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name        string
		strs        []string
		bfn         BoolFn
		expectedLen int
	}{
		{
			name: "Filters out strings that contain 'test'",
			strs: strSlice("test", "testing", "another"),
			bfn: func(s string, _ int, _ []string) bool {
				return !strings.Contains(s, "test")
			},
			expectedLen: 1,
		},
		{
			name: "Filters out empty strings",
			strs: strSlice("", "a string", "", "another string"),
			bfn: func(s string, _ int, _ []string) bool {
				return s != ""
			},
			expectedLen: 2,
		},
		{
			name: "Filters out len(string) > 10",
			strs: strSlice("a very long string"),
			bfn: func(s string, _ int, _ []string) bool {
				return len(s) < 10
			},
			expectedLen: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Filter(test.strs, test.bfn)

			assert.Equal(t, test.expectedLen, len(actual))
		})
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		bfn      BoolFn
		expected bool
	}{
		{
			name: "Slice contains string == slice",
			strs: strSlice("slice", "apple", "int"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "slice"
			},
			expected: true,
		},
		{
			name: "Slice does not string == slice",
			strs: strSlice("apple", "orange"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "slice"
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Some(test.strs, test.bfn)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		sfn      StringFn
		expected []string
	}{
		{
			name: "appends string to itself",
			strs: strSlice("apple", "orange"),
			sfn: func(s string, _ int, _ []string) string {
				return s + s
			},
			expected: strSlice("appleapple", "orangeorange"),
		},
		{
			name: "strips the final two letters",
			strs: strSlice("teletubbie", "telephone"),
			sfn: func(s string, _ int, _ []string) string {
				return s[:len(s)-2]
			},
			expected: strSlice("teletubb", "telepho"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Map(test.strs, test.sfn)

			for i := range actual {
				assert.Equal(t, test.expected[i], actual[i])
			}
		})
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		bfn      BoolFn
		expected bool
	}{
		{
			name: "Every len(string) > 2",
			strs: strSlice("greater than two", "123"),
			bfn: func(s string, _ int, _ []string) bool {
				return len(s) > 2
			},
			expected: true,
		},
		{
			name: "Every int not > 2",
			strs: strSlice("a", "beta", "capital"),
			bfn: func(s string, _ int, _ []string) bool {
				return len(s) == 100
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Every(test.strs, test.bfn)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		bfn      BoolFn
		expected func() *string
	}{
		{
			name: "Found returns int",
			strs: strSlice("string", "test", "ale"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "test"
			},
			expected: stringRef("test"),
		},
		{
			name: "Not found returns nil",
			strs: strSlice("test", "another"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "notfound"
			},
			expected: func() *string { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Find(test.strs, test.bfn)

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
		strs     []string
		bfn      BoolFn
		expected func() *int
	}{
		{
			name: "Found returns index of string",
			strs: strSlice("string", "another", "test"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "another"
			},
			expected: intRef(1),
		},
		{
			name: "Not found returns nil",
			strs: strSlice("test", "another"),
			bfn: func(s string, _ int, _ []string) bool {
				return s == "notfound"
			},
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := FindIndex(test.strs, test.bfn)

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
		strs     []string
		includes string
		expected bool
	}{
		{
			name:     "Found returns index of int",
			strs:     strSlice("testing", "test", "another"),
			includes: "test",
			expected: true,
		},
		{
			name:     "Not found returns nil",
			strs:     strSlice("womp", "chomp", "pomp"),
			includes: "lomp",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Includes(test.strs, test.includes)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		includes string
		expected func() *int
	}{
		{
			name:     "Found returns index of int",
			strs:     strSlice("notthisone", "testing", "foundit"),
			includes: "foundit",
			expected: intRef(2),
		},
		{
			name:     "Not found returns nil",
			strs:     strSlice("test", "testing"),
			includes: "notfound",
			expected: func() *int { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IndexOf(test.strs, test.includes)

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
		strs     []string
		join     string
		expected string
	}{
		{
			name:     "Join string on ,",
			strs:     strSlice("a", "b", "c"),
			join:     ",",
			expected: "a,b,c",
		},
		{
			name:     "Join string on /",
			strs:     strSlice("a", "b", "c"),
			join:     "/",
			expected: "a/b/c",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Join(test.strs, test.join)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected func() *string
	}{
		{
			name:     "Returns a ref to the last item in the list",
			strs:     strSlice("test", "another", "the last one"),
			expected: stringRef("the last one"),
		},
		{
			name:     "Empty slice returns nil",
			strs:     strSlice(),
			expected: func() *string { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Last(test.strs)

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
		strs     []string
		expected func() *string
	}{
		{
			name:     "Returns first int",
			strs:     strSlice("the first one", "test", "another"),
			expected: stringRef("the first one"),
		},
		{
			name:     "Empty list returns nil",
			strs:     strSlice(),
			expected: func() *string { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := First(test.strs)

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
		strs     []string
		at       int
		expected func() *string
	}{
		{
			name:     "Returns *int at index",
			strs:     strSlice("test", "another", "testing", "index 3"),
			at:       2,
			expected: stringRef("testing"),
		},
		{
			name:     "Empty list returns nil",
			strs:     strSlice(),
			at:       0,
			expected: func() *string { return nil },
		},
		{
			name:     "Negative returns nil",
			strs:     strSlice("testing", "another"),
			at:       -1,
			expected: func() *string { return nil },
		},
		{
			name:     "Index beyond length of slice",
			strs:     strSlice("testing", "another"),
			at:       5,
			expected: func() *string { return nil },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := At(test.strs, test.at)

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
		strs    []string
		prepend string
	}{
		{
			name:    "Prepends int",
			strs:    strSlice("test", "another"),
			prepend: "pre",
		},
		{
			name:    "Empty list returns list of 1 item",
			strs:    strSlice(),
			prepend: "pre",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Prepend(test.strs, test.prepend)

			assert.Equal(t, len(actual), len(test.strs)+1)
			assert.Equal(t, *actual.First(), test.prepend)
		})
	}
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected bool
	}{
		{
			name:     "Non-empty list",
			strs:     strSlice("testing"),
			expected: false,
		},
		{
			name:     "Empty list returns true",
			strs:     strSlice(),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Empty(test.strs)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected []string
	}{
		{
			name:     "Sorts slice",
			strs:     strSlice("alpha", "gamma", "delta", "beta"),
			expected: strSlice("alpha", "beta", "delta", "gamma"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Sort(test.strs)

			for i := range actual {
				assert.Equal(t, test.expected[i], actual[i])
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "Prepends int",
			strs: strSlice("test", "testing", "another"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := make([]string, 0, len(test.strs))
			actual := From(test.strs).Copy(expected)

			for i := range actual {
				assert.Equal(t, expected[i], actual[i])
			}
		})
	}
}

func TestBoolFn_And(t *testing.T) {
	tests := []struct {
		name     string
		first    BoolFn
		second   BoolFn
		expected bool
	}{
		{
			name:     "And composes two fns together",
			first:    boolFnComposer(true),
			second:   boolFnComposer(true),
			expected: true,
		},
		{
			name:     "Returns last false",
			first:    boolFnComposer(true),
			second:   boolFnComposer(false),
			expected: false,
		},
		{
			name:     "Returns first false",
			first:    boolFnComposer(false),
			second:   boolFnComposer(true),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.first.And(test.second)("", 0, []string{})

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestBoolFn_AndNot(t *testing.T) {
	tests := []struct {
		name     string
		first    BoolFn
		second   BoolFn
		expected bool
	}{
		{
			name:     "Returns true when Anded fn is returns false",
			first:    boolFnComposer(true),
			second:   boolFnComposer(false),
			expected: true,
		},
		{
			name:     "Returns false when Anded condition returns true",
			first:    boolFnComposer(true),
			second:   boolFnComposer(true),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.first.AndNot(test.second)("", 0, []string{})

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestBoolFn_Or(t *testing.T) {
	tests := []struct {
		name     string
		first    BoolFn
		second   BoolFn
		expected bool
	}{
		{
			name:     "And composes two fns together",
			first:    boolFnComposer(true),
			second:   boolFnComposer(true),
			expected: true,
		},
		{
			name:     "Returns true if either is true",
			first:    boolFnComposer(true),
			second:   boolFnComposer(false),
			expected: true,
		},
		{
			name:     "Returns false if both are false",
			first:    boolFnComposer(false),
			second:   boolFnComposer(false),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.first.Or(test.second)("", 0, []string{})

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestStringFn_And(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		first    StringFn
		second   StringFn
		expected string
	}{
		{
			name:     "And composes two fns together",
			initial:  "",
			first:    stringFnAccumulator("test"),
			second:   stringFnAccumulator("another"),
			expected: "testanother",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.first.And(test.second)(test.initial, 0, []string{})

			assert.Equal(t, test.expected, actual)
		})
	}
}
