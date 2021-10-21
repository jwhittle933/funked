package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	testmocks "github.com/jwhittle933/funked/testing"
)

// func addThree(t *Tree) {
// 	t.Add(NewBranch("first"), ModeReplace, "first").
// 		Add(NewBranch("second"), ModeReplace, "second").
// 		Add(NewBranch("third"), ModeReplace, "third")
// }

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
			tree := New(NewBranch(BranchRoot))
			tree.Add(nil, ModeReplace, test.addPath...)

			assert.Equal(t, test.expectedDepth, tree.depth)
		})
	}
}
