package binarytree_test

import (
	"testing"

	bt "github.com/ianfoo/ds-algo/binarytree"
)

func TestTraversal(t *testing.T) {
	tt := []struct {
		desc     string
		input    []int
		expected string
	}{
		{
			desc:     "in-order",
			input:    []int{12, 44, 21, 99, 143, 68, 69, 3, 14},
			expected: "3 12 14 21 44 68 69 99 143",
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			tree := bt.NewTree()
			for _, i := range tc.input {
				tree.Insert(i)
			}
			if o := tree.InOrder(); o != tc.expected {
				t.Errorf("expected %q, but got %q", tc.expected, o)
			}
		})
	}
}
