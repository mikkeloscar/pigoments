package pigoments

import (
	"fmt"
	"testing"

	"github.com/mikkeloscar/pigoments/formatters"
)

// func TestSome(t *testing.T) {
// 	code := `package formatters

// import (
// 	"bytes"
// 	"fmt"

// 	"github.com/mikkeloscar/pigoments/lexers"
// )

// var cssClasses = map[lexers.TokenType]string{
// 	lexers.TokenError: "err",
// 	// comments
// 	lexers.TokenComment:          "c",
// 	lexers.TokenCommentMultiline: "cm",
// 	lexers.TokenCommentPreproc:   "cp",
// 	lexers.TokenCommentSingle:    "c1",
// 	lexers.TokenCommentSpecial:   "cs",
// 	// generics
// 	lexers.TokenGenericDeleted:    "gd",
// 	lexers.TokenGenericEmph:       "ge",
// 	lexers.TokenGenericError:      "gr",
// 	lexers.TokenGenericHeading:    "gh",
// 	lexers.TokenGenericInserted:   "gi",
// 	lexers.TokenGenericOutput:     "go",
// 	lexers.TokenGenericPrompt:     "gp",
// 	lexers.TokenGenericStrong:     "gs",
// 	lexers.TokenGenericSubheading: "gu",
// 	lexers.TokenGenericTraceback:  "gt",
// 	// keyword
// 	lexers.TokenKeyword:            "k",
// 	lexers.TokenKeywordConstant:    "kc",
// 	lexers.TokenKeywordDeclaration: "kd",
// 	lexers.TokenKeywordNamespace:   "kn",
// 	lexers.TokenKeywordPseudo:      "kp",
// 	lexers.TokenKeywordReserved:    "kr",
// 	lexers.TokenKeywordType:        "kt",
// 	// literal
// 	lexers.TokenLiteralNumber:            "m",
// 	lexers.TokenLiteralString:            "s",
// 	lexers.TokenLiteralNumberFloat:       "mf",
// 	lexers.TokenLiteralNumberHex:         "mh",
// 	lexers.TokenLiteralNumberInteger:     "mi",
// 	lexers.TokenLiteralNumberOct:         "mo",
// 	lexers.TokenLiteralStringBacktick:    "sb",
// 	lexers.TokenLiteralStringChar:        "sc",
// 	lexers.TokenLiteralStringDoc:         "sd",
// 	lexers.TokenLiteralStringDouble:      "s2",
// 	lexers.TokenLiteralStringEscape:      "se",
// 	lexers.TokenLiteralStringHeredoc:     "sh",
// 	lexers.TokenLiteralStringInterpol:    "si",
// 	lexers.TokenLiteralStringOther:       "sx",
// 	lexers.TokenLiteralStringRegex:       "sr",
// 	lexers.TokenLiteralStringSingle:      "s1",
// 	lexers.TokenLiteralStringSymbol:      "ss",
// 	lexers.TokenLiteralNumberIntegerLong: "il",
// 	// name
// 	lexers.TokenNameAttribute:        "na",
// 	lexers.TokenNameBuiltin:          "nb",
// 	lexers.TokenNameClass:            "nc",
// 	lexers.TokenNameConstant:         "no",
// 	lexers.TokenNameDecorator:        "nd",
// 	lexers.TokenNameEntity:           "ni",
// 	lexers.TokenNameException:        "ne",
// 	lexers.TokenNameFunction:         "nf",
// 	lexers.TokenNameLabel:            "nl",
// 	lexers.TokenNameNamespace:        "nn",
// 	lexers.TokenNameTag:              "nt",
// 	lexers.TokenName:                 "nx",
// 	lexers.TokenNameVariable:         "nv",
// 	lexers.TokenNameBuiltinPseudo:    "bp",
// 	lexers.TokenNameVariableClass:    "vc",
// 	lexers.TokenNameVariableGlobal:   "vg",
// 	lexers.TokenNameVariableInstance: "vi",
// 	// operator
// 	lexers.TokenOperator:     "o",
// 	lexers.TokenOperatorWord: "ow",
// 	// text
// 	lexers.TokenTextWhitespace: "w",
// 	// punctuation
// 	lexers.TokenPunctuation: "p",
// }

// var escape_html_table = map[rune]string{
// 	'&':  "&amp;",
// 	'<':  "&lt;",
// 	'>':  "&gt;",
// 	'"':  "&quot;",
// 	'\'': "&#39;",
// }

// const wrapper = "<div class=\"%s\"><pre>\n%s\n</pre></div>"

// // HTMLFormatter can create html formatted output
// type HTMLFormatter struct {
// 	CssClass string
// }

// // Generate can generate the html
// func (f *HTMLFormatter) Generate(l *lexers.Lexer) string {
// 	var buf bytes.Buffer
// 	var token *lexers.Token
// 	for {
// 		token = l.NextToken()
// 		if token == nil {
// 			return fmt.Sprintf(wrapper, f.CssClass, buf.String())
// 		}

// 		if css, ok := cssClasses[token.Type]; ok {
// 			buf.WriteString(fmt.Sprintf(` + "`" + `<span class="%s">` + "`" + `, css))
// 			buf.WriteString(escape(token.Value))
// 			buf.WriteString("</span>")
// 		} else {
// 			buf.WriteString(token.Value)
// 		}
// 	}
// }

// func escape(str string) string {
// 	var buf bytes.Buffer

// 	for _, c := range str {
// 		if e, ok := escape_html_table[c]; ok {
// 			buf.WriteString(e)
// 		} else {
// 			buf.WriteRune(c)
// 		}
// 	}

// 	return buf.String()
// }`
// 	// fmt.Println(code)
// 	f := &formatters.HTMLFormatter{CssClass: "highlight"}
// 	_, err := Highlight(code, "go", f)
// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}
// 	// fmt.Printf("%s\n", h)
// }

// func TestIni(t *testing.T) {
// 	code := `# comment

// [main]
// key = value
// ; comment 2`

// 	f := &formatters.HTMLFormatter{LineNos: true}
// 	out, err := Highlight(code, "ini", f)
// 	if err != nil {
// 		t.Errorf("%s", err.Error())
// 	}
// 	fmt.Println(out)
// }

func TestHTML(t *testing.T) {
	code := `<html><head>
<style>
.css {
	font-weight: bold;
}
</style>
<body>
<!-- body -->
<div class="body"></div>
</body>
</html>`

	code = `package lexers

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

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

const eof = -1

type Pos int

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// lexer holds the state of the scanner.
// TODO comments
type Lexer struct {
	input   string      // the string being scanned
	state   stateFn     // the next lexing function to enter
	pos     Pos         // current position in the input
	start   Pos         // start position of this item
	width   Pos         // width of last rune read from input
	lastPos Pos         // position of most recent item returned by nextItem
	tokens  chan *Token // channel of scanned items
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *Lexer) emit(t TokenType) {
	l.tokens <- &Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) token() string {
	return l.input[l.start:l.pos]
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
func (l *Lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.lastPos], "\n")
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- &Token{TokenError, fmt.Sprintf(format, args...)}
	return nil
}

// nextItem returns the next item from the input.
func (l *Lexer) NextToken() *Token {
	token := <-l.tokens
	// TODO l.lastPos = item.pos
	return token
}

// lex creates a new scanner for the input string.
func Lex(input, lang string) *Lexer {
	l := &Lexer{
		input:  input,
		tokens: make(chan *Token),
	}
	go l.run(lexText)
	return l
}

// run runs the state machine for the lexer.
func (l *Lexer) run(startState stateFn) {
	for l.state = startState; l.state != nil; {
		l.state = l.state(l)
	}
	l.tokens <- nil // inform listener that we are done
}

// lexText scans until an opening action delimiter, "{{".
func lexText(l *Lexer) stateFn {
	fmt.Println("hej")
	switch p := l.peek(); {
	case isAlpha(p):
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
	case p == '` + "`" + `':
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
		case !isAlpha(c):
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
		case '` + "`" + `':
			break Loop
		}
	}
	l.emit(TokenLiteralString)
	return lexText
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isWhiteSpace(r rune) bool {
	return strings.IndexRune(" \t\n\r\f", r) >= 0
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

func isAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || unicode.IsDigit(r)
}

func isPunctuation(r rune) bool {
	return strings.IndexRune("|^<>=!()[]{}.,;:", r) >= 0
}`

	f := &formatters.HTMLFormatter{LineNos: true}
	out, err := Highlight(code, "html", f)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	fmt.Println(out)
}
