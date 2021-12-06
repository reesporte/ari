package lex

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIsOp(t *testing.T) {
	if !isOp("+") {
		t.Errorf("+ is an operation")
	}
	if !isOp("-") {
		t.Errorf("- is an operation")
	}
	if !isOp("*") {
		t.Errorf("* is an operation")
	}
	if !isOp("/") {
		t.Errorf("/ is an operation")
	}
	if isOp("f") {
		t.Errorf("f is not an operation")
	}
	if isOp("0") {
		t.Errorf("0 is not an operation")
	}
}

func TestIsNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		if !isNum(strconv.Itoa(i)) {
			t.Fatalf("%v is a numeric", i)
		}
	}
	if isNum("applesauce") {
		t.Fatal("applesauce is not a numeric")
	}
}

func TestLex(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1 * 6",
			expected: []Token{
				Token{NUM, "1"}, Token{OP, "*"}, Token{NUM, "6"}},
		},
		{
			input: "7 / 65",
			expected: []Token{
				Token{NUM, "7"}, Token{OP, "/"}, Token{NUM, "65"}},
		},
		{
			input: "0 + 0",
			expected: []Token{
				Token{NUM, "0"}, Token{OP, "+"}, Token{NUM, "0"}},
		},
		{
			input: "0 + 0 + 5 * 3 - 2 + 1",
			expected: []Token{
				Token{NUM, "0"}, Token{OP, "+"}, Token{NUM, "0"}, Token{OP, "+"}, Token{NUM, "5"}, Token{OP, "*"}, Token{NUM, "3"}, Token{OP, "-"}, Token{NUM, "2"}, Token{OP, "+"}, Token{NUM, "1"}},
		},
	}
	for _, test := range tests {
		result := Lex(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, result)
		}
	}
}
