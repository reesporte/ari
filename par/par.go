package par

import (
	"ari/lex"
	"fmt"
	"strconv"
)

// Stack is a basic stack
type Stack []*Node

// IsEmpty returns whether the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds val to the stack
func (s *Stack) Push(val *Node) {
	*s = append(*s, val)
}

// Pop removes an element from the stack and returns it
func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element
}

// Peek returns the top element of the stack but doesn't remove it
func (s *Stack) Peek() *Node {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

type Node struct {
	left  *Node
	right *Node
	op    string
	num   float64
}

// height gets the height of a given tree with root node `n`
func (n Node) height() int {
	lheight := 0
	rheight := 0
	if n.left != nil {
		lheight = n.left.height()
	}
	if n.right != nil {
		rheight = n.right.height()
	}
	if lheight > rheight {
		return lheight + 1
	}
	return rheight + 1
}

// PrintTree prints the tree in breadth-first order
func (n Node) PrintTree() {
	h := n.height()
	for i := 1; i < h+1; i++ {
		n.printLevel(i)
		fmt.Println()
	}
}

// printLevel prints the data at a given level of a tree
func (n Node) printLevel(level int) {
	if level == 1 {
		if n.op != "" {
			fmt.Printf("%v ", n.op)
		} else {
			fmt.Printf("%v ", n.num)
		}
	} else if level > 1 {
		if n.left != nil {
			n.left.printLevel(level - 1)
		}
		if n.right != nil {
			n.right.printLevel(level - 1)
		}
	}
}

func (n *Node) Interpret() float64 {
	return n.eval()
}

func (n *Node) eval() float64 {
	if n == nil {
		return 0
	}
	if n.op == "" {
		return n.num
	}
	if n.left != nil && n.right != nil {
		left := n.left.eval()
		right := n.right.eval()
		switch n.op {
		case "+":
			return left + right
		case "-":
			return left - right
		case "/":
			return left / right
		case "*":
			return left * right
		}
	}
	panic("invalid expression parsed!")
}

// Parse takes in a line of text and parses it into an AST
func Parse(line string) Node {
	tkns := lex.Lex(line)
	operators := Stack{}
	operands := Stack{}
	for _, tkn := range tkns {
		if tkn.Class == lex.NUM {
			val, _ := strconv.ParseFloat(tkn.Repr, 64)
			operands.Push(&Node{num: val})
		} else if tkn.Class == lex.OP {
			t := &Node{op: tkn.Repr}
			for operators.Peek() != nil && greaterPrecedence(operators.Peek(), t) {
				op := operators.Pop()
				op.right = operands.Pop()
				op.left = operands.Pop()
				operands.Push(op)
			}
			operators.Push(t)
		}
	}
	for operators.Peek() != nil {
		op := operators.Pop()
		op.right = operands.Pop()
		op.left = operands.Pop()
		operands.Push(op)
	}
	return *operands.Pop()
}

var precedences = map[string]int{
	"/": 2,
	"*": 1,
	"-": 0,
	"+": 0,
}

func greaterPrecedence(a, b *Node) bool {
	if a.op == "" {
		return false
	}
	return precedences[a.op] > precedences[b.op]
}
