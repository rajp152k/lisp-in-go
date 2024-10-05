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
		"add": lextest{"(+ 1 2)",
			[]Token{
				*getSpecial["("],
				Token{"S", "+"},
				Token{"S", "1"},
				Token{"S", "2"},
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
			if !equalTokenStream(lexed, lext.lexed) {
				t.Fail()
			}

		})

	}
}
