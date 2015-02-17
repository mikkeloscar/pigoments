package lexers

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var lexers = map[string]stateFn{
	"go":   goLexer,
	"diff": diffLexer,
	"html": htmlLexer,
}

const eof = -1

// Pos defines a position in the input
type Pos int

// stateFn represents the state of the lexer as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// stack of states in the lexer
type stateStack []stateFn

// get top state of stack
func (s stateStack) top() stateFn {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return lexText
}

// pop state of stack
func (s *stateStack) pop() stateFn {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

// push state onto stack
func (s *stateStack) push(st stateFn) {
	(*s) = append((*s), st)
}

// Lexer holds the state of the scanner.
type Lexer struct {
	input   string  // the string being scanned
	state   stateFn // the next lexing function to enter
	stack   stateStack
	escape  *regexp.Regexp
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
	if l.start < l.pos {
		l.tokens <- &Token{t, l.input[l.start:l.pos]}
		l.start = l.pos
	}
}

// get current token
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

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() *Token {
	token := <-l.tokens
	// TODO l.lastPos = item.pos
	return token
}

// Lex creates a new scanner for the input string.
func Lex(input, lang string) *Lexer {
	l := &Lexer{
		input:  input,
		tokens: make(chan *Token),
	}
	// add lexer to stack
	l.stack.push(lexers[lang])
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *Lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
	l.tokens <- nil // inform listener that we are done
}

// default lexer
func lexText(l *Lexer) stateFn {
	// if we are in a sub state check if we should leave it
	if l.escape != nil && l.escape.MatchString(l.input[l.start:]) {
		l.stack.pop()
		l.escape = nil
	}

	return l.stack.top()
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isWhiteSpace reports whether r is a whitespace character.
func isWhiteSpace(r rune) bool {
	return strings.IndexRune(" \t\n\r\f", r) >= 0
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isUAlpha reports whether r is a unicode alpha character.
func isUAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isUAlphaNumeric(r rune) bool {
	return isUAlpha(r) || unicode.IsDigit(r)
}
