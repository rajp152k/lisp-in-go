package lexer

import (
	"testing"
)

type lextest struct {
	source string
	lexed  []Token
}

func yieldLexTests() map[string]lextest {
	return map[string]lextest{
		"empty": lextest{"",
			[]Token{},
		},
		"nil": lextest{"'()",
			[]Token{
				*getSpecial["'"],
				*getSpecial["("],
				*getSpecial[")"],
			},
		},
		"symbol": lextest{"symbol",
			[]Token{
				Token{"S", "symbol"},
			},
		},
		"add": lextest{"(+ 1 2)",
			[]Token{
				*getSpecial["("],
				Token{"S", "+"},
				*getSpecial[" "],
				Token{"S", "1"},
				*getSpecial[" "],
				Token{"S", "2"},
				*getSpecial[")"],
			},
		},
		"funcall": lextest{"(square x)",
			[]Token{
				*getSpecial["("],
				Token{"S", "square"},
				*getSpecial[" "],
				Token{"S", "x"},
				*getSpecial[")"],
			},
		},
		"nested": lextest{"(lambda (x y) x)",
			[]Token{
				*getSpecial["("],
				Token{"S", "lambda"},
				*getSpecial[" "],
				*getSpecial["("],
				Token{"S", "x"},
				*getSpecial[" "],
				Token{"S", "y"},
				*getSpecial[")"],
				*getSpecial[" "],
				Token{"S", "x"},
				*getSpecial[")"],
			},
		},
		"string": lextest{"|this is a string|",
			[]Token{
				Token{"STR", "this is a string"},
			},
		},
		"concat": lextest{"(concat |this is a string| |this is another string|)",
			[]Token{
				*getSpecial["("],
				Token{"S", "concat"},
				*getSpecial[" "],
				Token{"STR", "this is a string"},
				*getSpecial[" "],
				Token{"STR", "this is another string"},
				*getSpecial[")"],
			},
		},
	}
}

func TestLexer(t *testing.T) {
	for name, lext := range yieldLexTests() {
		name, lext := name, lext
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			lexed := Lexer(lext.source)
			l.Println(lexed)
			if !equalTokenStream(lexed, lext.lexed) {
				t.Fail()
			}

		})
	}
}
