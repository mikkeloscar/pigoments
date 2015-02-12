package lexers

// diff lexer
func diffLexer(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '\n':
			if len(l.token()) > 1 {
				switch l.token()[0] {
				case ' ':
					l.emit(TokenText)
				case '+':
					l.emit(TokenGenericInserted)
				case '-':
					l.emit(TokenGenericDeleted)
				case '!':
					l.emit(TokenGenericStrong)
				case '=':
					l.emit(TokenGenericHeading)
				case 'i', 'I':
					if len(l.token()) > 5 && l.token()[1:5] == "ndex" {
						l.emit(TokenGenericHeading)
					} else {
						l.emit(TokenText)
					}
				case 'd':
					if len(l.token()) > 4 && l.token()[1:4] == "iff" {
						l.emit(TokenGenericHeading)
					} else {
						l.emit(TokenText)
					}
				default:
					l.emit(TokenText)
				}
			} else {
				l.emit(TokenText)
			}
		case '@':
			if l.next() == '@' {
				return diffLexPosition
			}
			l.backup()
		case eof:
			l.emit(TokenText)
			return nil
		}
	}
}

func diffLexPosition(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '@':
			if l.next() == '@' {
				l.emit(TokenGenericSubheading)
			}
		case '\n', eof:
			l.emit(TokenText)
			return lexText
		}
	}
}
