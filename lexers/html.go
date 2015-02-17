package lexers

import (
	"regexp"
	"strings"
)

var htmlScriptEscape = regexp.MustCompile(`^<\s*/\s*script\s*>`)
var htmlStyleEscape = regexp.MustCompile(`^<\s*/\s*style\s*>`)

// html lexer
func htmlLexer(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '<':
			l.backup()
			l.emit(TokenText)
			l.next()
			return htmlLexStartBracket
		case eof:
			l.emit(TokenText)
			return nil
		}
	}
}

func htmlLexStartBracket(l *Lexer) stateFn {
	l.acceptRun(" \t")
	switch l.next() {
	case '!':
		if l.next() == '-' && l.next() == '-' {
			return htmlLexComment
		}
		return htmlLexCommentPreproc
	case '/':
		l.emit(TokenPunctuation)
		return htmlLexTag
	default:
		l.backup()
		l.emit(TokenPunctuation)
		l.next()
		return htmlLexTag
	}
}

func htmlLexComment(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '-':
			if l.next() == '-' && l.next() == '>' {
				l.emit(TokenComment)
				return lexText
			}
		case eof:
			l.emit(TokenComment) // hit an unclosed comment
			return nil
		}
	}
}

func htmlLexCommentPreproc(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '>':
			l.emit(TokenCommentPreproc)
			return lexText
		case eof:
			l.emit(TokenCommentPreproc)
			return nil
		}
	}
}

func htmlLexTag(l *Lexer) stateFn {
	for {
		switch l.next() {
		case ' ', '>':
			l.backup()
			if l.token() == "script" {
				// TODO l.stack.push(jsLexer)
				l.escape = htmlScriptEscape
			}

			if l.token() == "style" {
				// TODO l.stack.push(cssLexer)
				l.escape = htmlStyleEscape
			}

			l.emit(TokenNameTag)
			return htmlLexAttribute
		}
	}
	return lexText
}

func htmlLexAttribute(l *Lexer) stateFn {
	for {
		switch l.next() {
		case '=':
			l.backup()
			l.emit(TokenNameAttribute)
			l.next()
			l.emit(TokenPunctuation) // TODO correct?
		case '"':
			return htmlLexDoubleAttrVal
		case '\'':
			return htmlLexSingleAttrVal
		case '>':
			l.emit(TokenPunctuation)
			return lexText
		case eof:
			l.emit(TokenText)
			return nil
		}
	}
}

func htmlLexSingleAttrVal(l *Lexer) stateFn {
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
	l.emit(TokenLiteralString)
	return htmlLexAttribute
}

func htmlLexDoubleAttrVal(l *Lexer) stateFn {
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
	return htmlLexAttribute
}

// isHTMLPunctuation reports whether r is a special html character.
func isHTMLPunctuation(r rune) bool {
	return strings.IndexRune("<&", r) >= 0
}
