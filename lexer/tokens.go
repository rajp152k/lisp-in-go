package lexer

type Token struct {
	id   string
	repr string
}

type Symbol struct {
	Token
}

type Special struct {
	Token
}

func newToken(id string, repr string) *Token {
	return &Token{id, repr}
}

func newSymbol(repr string) *Symbol {
	return &Symbol{*newToken("S", repr)}
}

type TokenStream []Token

var getSpecial map[string]*Token

func init() {
	//initializing special tokens
	getSpecial = map[string](*Token){}
	getSpecial["("] = newToken("LP", "(")
	getSpecial[")"] = newToken("RP", ")")
	getSpecial["'"] = newToken("Q", "'")
	getSpecial["`"] = newToken("QQ", "`")
	getSpecial[","] = newToken("UQ", ",")
}
