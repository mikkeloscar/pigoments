package lexers

import "unicode"

// keyword maps
var namespaceKeywords = map[string]struct{}{
	"import":  struct{}{},
	"package": struct{}{},
}

var declarationKeywords = map[string]struct{}{
	"var":       struct{}{},
	"func":      struct{}{},
	"struct":    struct{}{},
	"map":       struct{}{},
	"chan":      struct{}{},
	"type":      struct{}{},
	"interface": struct{}{},
	"const":     struct{}{},
}

var keywords = map[string]struct{}{
	"break":       struct{}{},
	"default":     struct{}{},
	"select":      struct{}{},
	"case":        struct{}{},
	"defer":       struct{}{},
	"go":          struct{}{},
	"else":        struct{}{},
	"goto":        struct{}{},
	"switch":      struct{}{},
	"fallthrough": struct{}{},
	"if":          struct{}{},
	"range":       struct{}{},
	"continue":    struct{}{},
	"for":         struct{}{},
	"return":      struct{}{},
}

var constants = map[string]struct{}{
	"true":  struct{}{},
	"false": struct{}{},
	"iota":  struct{}{},
	"nil":   struct{}{},
}

var builtins = map[string]struct{}{
	"uint":       struct{}{},
	"uint8":      struct{}{},
	"uint16":     struct{}{},
	"uint32":     struct{}{},
	"uint64":     struct{}{},
	"int":        struct{}{},
	"int8":       struct{}{},
	"int16":      struct{}{},
	"int32":      struct{}{},
	"int64":      struct{}{},
	"float":      struct{}{},
	"float32":    struct{}{},
	"float64":    struct{}{},
	"complex64":  struct{}{},
	"complex128": struct{}{},
	"byte":       struct{}{},
	"rune":       struct{}{},
	"string":     struct{}{},
	"bool":       struct{}{},
	"error":      struct{}{},
	"uintptr":    struct{}{},
	"print":      struct{}{},
	"println":    struct{}{},
	"panic":      struct{}{},
	"recover":    struct{}{},
	"close":      struct{}{},
	"complex":    struct{}{},
	"real":       struct{}{},
	"imag":       struct{}{},
	"len":        struct{}{},
	"cap":        struct{}{},
	"append":     struct{}{},
	"copy":       struct{}{},
	"delete":     struct{}{},
	"new":        struct{}{},
	"make":       struct{}{},
}

var types = map[string]struct{}{
	"uint":       struct{}{},
	"uint8":      struct{}{},
	"uint16":     struct{}{},
	"uint32":     struct{}{},
	"uint64":     struct{}{},
	"int":        struct{}{},
	"int8":       struct{}{},
	"int16":      struct{}{},
	"int32":      struct{}{},
	"int64":      struct{}{},
	"float":      struct{}{},
	"float32":    struct{}{},
	"float64":    struct{}{},
	"complex64":  struct{}{},
	"complex128": struct{}{},
	"byte":       struct{}{},
	"rune":       struct{}{},
	"string":     struct{}{},
	"bool":       struct{}{},
	"error":      struct{}{},
	"uintptr":    struct{}{},
}

// entry lexer refer to this with lexText
func goLexer(l *Lexer) stateFn {
	switch p := l.peek(); {
	case isUAlpha(p):
		l.next()
		return lexKeywordIdentifier
	case isWhiteSpace(p):
		return lexWhitespace
	case p == '/':
		l.next()
		if l.peek() == '/' {
			l.next()
			return lexSingleComment
		}
		l.emit(TokenText)
		return lexText
	case unicode.IsDigit(p):
		return lexNumber
	case p == '.':
		l.next()
		if unicode.IsDigit(l.peek()) {
			l.backup()
			return lexNumber
		}
		// TODO
		l.emit(TokenPunctuation)
		return lexText
	case p == '"':
		l.next()
		return lexString
	case p == '`':
		l.next()
		return lexRawString
	case isPunctuation(p):
		l.next()
		l.emit(TokenPunctuation)
		return lexText
	case p == eof:
		return nil
	default:
		l.next()
		l.emit(TokenText)
		return lexText
	}
}

func lexWhitespace(l *Lexer) stateFn {
	for {
		switch l.next() {
		case ' ', '\t', '\n', '\r', '\f':
			// absorb
		default:
			l.backup()
			l.emit(TokenText)
			return lexText
		}
	}
}

func lexSingleComment(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '\n', '\r':
			l.backup()
			l.emit(TokenCommentSingle)
			l.next()
			l.emit(TokenText)
			return lexText
		}
	}
}

func lexKeywordIdentifier(l *Lexer) stateFn {
	for {
		switch c := l.next(); {
		case !isUAlpha(c):
			l.backup()
			if _, ok := namespaceKeywords[l.token()]; ok {
				l.emit(TokenKeywordNamespace)
			} else if _, ok := declarationKeywords[l.token()]; ok {
				l.emit(TokenKeywordDeclaration)
			} else if _, ok := keywords[l.token()]; ok {
				l.emit(TokenKeyword)
			} else if _, ok := constants[l.token()]; ok {
				l.emit(TokenKeywordConstant)
			} else if _, ok := builtins[l.token()]; ok {
				if l.peek() == '(' {
					l.emit(TokenNameBuiltin)
					l.next()
					l.emit(TokenPunctuation)
				}
			} else if _, ok := types[l.token()]; ok {
				l.emit(TokenKeywordType)
			} else {
				l.emit(TokenName)
			}
			return lexText
		}
	}
}

func lexNumber(l *Lexer) stateFn {
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		l.acceptRun("0123456789abcdefABCDEF")
		l.emit(TokenLiteralNumberHex)
		return lexText
	} else if l.accept("0") && l.accept("01234567") {
		l.acceptRun("01234567")
		if l.accept("i") {
			l.emit(TokenLiteralNumber)
		} else {
			l.emit(TokenLiteralNumberOct)
		}
		return lexText
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	if l.accept("i") {
		l.emit(TokenLiteralNumber)
	} else {
		l.emit(TokenLiteralNumberFloat)
	}
	// if isAlphaNumeric(l.peek()) {
	// 	l.next()
	// 	return false
	// }
	return lexText
}

// lexChar scans a character constant. The initial quote is already
// scanned. Syntax checking is done by the parser.
// TODO invalid char length
func lexChar(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case '\n':
			l.emit(TokenText)
			return lexText
		case eof:
			return nil
		case '\'':
			break Loop
		}
	}
	l.emit(TokenLiteralStringChar)
	return lexText
}

func lexString(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case '\n':
			l.emit(TokenText)
			return lexText
		case eof:
			return nil
		case '"':
			break Loop
		}
	}
	l.emit(TokenLiteralString)
	return lexText
}

func lexRawString(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof:
			return nil
		case '`':
			break Loop
		}
	}
	l.emit(TokenLiteralString)
	return lexText
}
