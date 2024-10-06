package lexer


func Lexer(source string) []Token {
	lexed := []Token{}
	var acc string
	var curr string

	resetAcc := func() {
		if acc != "" {
			lexed = append(lexed, *newSymbol(acc))
			acc = ""
		}
	}


	for i := 0; i < len(source); i++ {
		curr = string(source[i])
		if curr == "|"{
			resetAcc()
			for j := i+1; j < len(source); j++{
				currj := string(source[j])
				if currj != "|"{
					acc += currj
				} else {
					lexed = append(lexed, *newString(acc))
					acc = ""
					i = j
					break
				}
			}
			continue
		}
		if isSpecial(curr) {
			resetAcc()
			lexed = append(lexed, *getSpecial[curr])
		} else {
			acc += curr
		}
	}
	resetAcc()

	return lexed
}
