package lexers

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const eof = -1

type Pos int

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// lexer holds the state of the scanner.
// TODO comments
type Lexer struct {
	input   string      // the string being scanned
	state   stateFn     // the next lexing function to enter
	lang    stateFn     // current lexer
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
		input: input,
		// TODO use lang argument to decide this
		lang:   goLexer,
		tokens: make(chan *Token),
	}
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

func lexText(l *Lexer) stateFn {
	// TODO make stack
	return l.lang
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

func isUAlpha(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isUAlphaNumeric(r rune) bool {
	return isUAlpha(r) || unicode.IsDigit(r)
}

func isPunctuation(r rune) bool {
	return strings.IndexRune("|^<>=!()[]{}.,;:", r) >= 0
}
