package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	testmocks "github.com/jwhittle933/funked/tests"
)

func addThree(t *Tree) {
	t.Add(New("first"), ModeReplace.Uint8(), "first")
	t.Add(New("second"), ModeReplace.Uint8(), "first", "second")
	t.Add(New("third"), ModeReplace.Uint8(), "first", "second", "third")
}

func TestTree_Add(t *testing.T) {
	tests := []struct {
		addPath       []string
		expectedDepth int
		testmocks.TestCase
	}{
		{
			TestCase:      testmocks.TestCase{Name: "Empty path, tree depth 0"},
			expectedDepth: 0,
		},
		{
			TestCase:      testmocks.TestCase{Name: "Path of 3, tree depth of 3"},
			addPath:       []string{"first", "second", "third"},
			expectedDepth: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tr := New(RootNode)
			tr.Add(nil, ModeReplace.Uint8(), test.addPath...)

			assert.Equal(t, test.expectedDepth, tr.Depth())
		})
	}
}

func TestTree_Find(t *testing.T) {
	tests := []struct {
		testmocks.TestCase
		expectedDepth int
		nodeName      string
		nilNode       bool
		setup         func(t *Tree) (findPath []string)
	}{
		{
			TestCase:      testmocks.TestCase{Name: "Empty path should return depth of 0 and nil"},
			expectedDepth: 0,
			nodeName:      "root",
			nilNode:       true,
			setup: func(t *Tree) (findPath []string) {
				addThree(t)

				return []string{}
			},
		},
		{
			TestCase:      testmocks.TestCase{Name: "Three nodes should return depth of 3"},
			expectedDepth: 3,
			nodeName:      "third",
			setup: func(t *Tree) (findPath []string) {
				addThree(t)

				return []string{"first", "second", "third"}
			},
		},
		{
			TestCase:      testmocks.TestCase{Name: "Find second of 3 values"},
			expectedDepth: 2,
			nodeName:      "second",
			setup: func(t *Tree) (findPath []string) {
				addThree(t)

				return []string{"first", "second"}
			},
		},
		{
			TestCase:      testmocks.TestCase{Name: "Returns the last found node of a bad path"},
			expectedDepth: 3,
			nodeName:      "third",
			setup: func(t *Tree) (findPath []string) {
				addThree(t)

				return []string{"first", "second", "third", "fourth"}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tr := New(RootNode)
			findPath := test.setup(tr)
			depth, n := tr.Find(findPath...)

			assert.Equal(t, test.expectedDepth, depth)
			assert.True(t, IsNilNode(n) == test.nilNode)

			if !IsNilNode(n) {
				assert.Equal(t, test.nodeName, n.Name())
			}
		})
	}
}
