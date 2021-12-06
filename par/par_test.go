package par

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTreeHeight(t *testing.T) {
	root := Node{op: "+"}
	root.left = &Node{num: 1}
	root.right = &Node{op: "*"}
	root.right.left = &Node{num: 3}
	root.right.right = &Node{op: "*"}
	root.right.right.right = &Node{num: 3}
	root.right.right.left = &Node{num: 3}
	h := root.height()
	if h != 4 {
		t.Fatalf("tree height incorrect. expected 4, got %v", h)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		expected Node
		input    string
	}{
		{
			expected: Node{
				op:   "+",
				left: &Node{num: 1},
				right: &Node{
					op:    "+",
					left:  &Node{num: 1},
					right: &Node{num: 1},
				},
			},
			input: "1 + 1 + 1",
		},
		{
			expected: Node{
				op:   "-",
				left: &Node{num: 69},
				right: &Node{
					op:    "*",
					left:  &Node{num: 4},
					right: &Node{num: 5},
				},
			},
			input: "69 - 4 * 5",
		},
		{
			expected: Node{
				op:    "*",
				right: &Node{num: 5},
				left: &Node{
					op:    "/",
					left:  &Node{num: 69},
					right: &Node{num: 4},
				},
			},
			input: "69 / 4 * 5",
		},
		{
			expected: Node{
				op: "-",
				left: &Node{
					op:    "/",
					left:  &Node{num: 69},
					right: &Node{num: 4},
				},
				right: &Node{num: 5},
			},
			input: "69 / 4 - 5",
		},
		{
			expected: Node{
				op: "+",
				left: &Node{
					op:    "*",
					left:  &Node{num: 69},
					right: &Node{num: 4},
				},
				right: &Node{num: 5},
			},
			input: "69 * 4 + 5",
		},
		{
			expected: Node{num: 1},
			input:    "1",
		},
		{
			expected: Node{op: "+",
				left:  &Node{num: 420},
				right: &Node{num: 69},
			},
			input: "420 + 69",
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			n := Parse(test.input)
			if !reflect.DeepEqual(n, test.expected) {
				fmt.Println("expected:")
				test.expected.PrintTree()
				fmt.Println("got:")
				n.PrintTree()
				t.Error("invalid parsing")
			}
		})
	}
}

func TestInterpret(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "1 + 1 + 1",
			expected: 3.0,
		},
		{
			input:    "69 * 4 + 5",
			expected: 281.0,
		},
		{
			input:    "69 / 4 - 5",
			expected: 12.25,
		},
		{
			input:    "420 + 69 * 6969 / 3000.4321",
			expected: 580.2639166538713,
		},
		{
			input:    "0.420 + 0.69",
			expected: 1.1099999999999999,
		},
		{
			input:    ".42 + .6",
			expected: 1.02,
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ast := Parse(test.input)
			result := ast.Interpret()
			if result != test.expected {
				t.Errorf("expected %v got %v", test.expected, result)
			}
		})
	}
}
