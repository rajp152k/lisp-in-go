package lexer

import (
	"strings"
)

type Token struct {
	id   string
	repr string
}

func newToken(id string, repr string) *Token {
	return &Token{id, repr}
}

func newSymbol(repr string) *Token {
	return newToken("S", repr)
}

var getSpecial map[string]*Token

func isSpecial(id string) bool {
	return strings.Contains("()'`;\t \n", id)
}

func init() {
	//initializing special tokens
	getSpecial = map[string](*Token){}
	getSpecial["("] = newToken("LP", "(")
	getSpecial[")"] = newToken("RP", ")")
	getSpecial["'"] = newToken("Q", "'")
	getSpecial["`"] = newToken("QQ", "`")
	getSpecial[","] = newToken("UQ", ",")
	getSpecial["\t"] = newToken("TB", "\t")
	getSpecial[" "] = newToken("SP", " ")
	getSpecial["\n"] = newToken("NL", "\n")
}
