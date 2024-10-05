package lexer

import (
	"log"
	"os"
	"reflect"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "[lexer]: ", log.LstdFlags)
}

func equalTokenStream(t1, t2 []Token) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if len(t1) != len(t2) {
		return false
	}

	for i := range t1 {
		if !reflect.DeepEqual(t1[i], t2[i]) {
			return false
		}
	}
	return true
}
