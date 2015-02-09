package formatters

import "github.com/mikkeloscar/pigoments/lexers"

// Formatter interface
type Formatter interface {
	Generate(*lexers.Lexer) string
}
