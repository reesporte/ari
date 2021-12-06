// lex
package lex

import (
	"strings"
)

type TokenType int

const (
	NUM TokenType = iota
	OP
)

func (t TokenType) String() string {
	return [...]string{"NUM", "OP"}[t]
}

// Token holds a token's class and representation
type Token struct {
	Class TokenType
	Repr  string
}

// checks if a String is an operator
func isOp(tkn string) bool {
	return tkn == "+" || tkn == "-" || tkn == "/" || tkn == "*"
}

// checks if a string is a digit
func isNum(tkn string) bool {
	for _, r := range tkn {
		if !('0' <= r && r <= '9') {
			return false
		}
	}
	return true
}

// Lex lexes a given line for tokens
func Lex(line string) []Token {
	tkns := make([]Token, 0)
	for _, tkn := range strings.Split(line, " ") {
		if strings.TrimSpace(tkn) == "" {
			continue
		}
		// ignore anything that isnt an operator or a number
		var class TokenType
		if isOp(tkn) {
			class = OP
		} else if isNum(tkn) {
			class = NUM
		}

		tkns = append(tkns, Token{class, tkn})
	}
	return tkns
}
