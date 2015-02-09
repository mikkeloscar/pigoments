package pigoments

import (
	"fmt"

	"github.com/mikkeloscar/pigoments/formatters"
	"github.com/mikkeloscar/pigoments/lexers"
)

// Highlight code as language lang and return the formatted output by the
// formatter
func Highlight(code, lang string, formatter formatters.Formatter) (string, error) {
	l := lexers.Lex(code, lang)
	if l == nil {
		return "", fmt.Errorf("language: %s not supported", lang)
	}

	formatted := formatter.Generate(l)
	return formatted, nil
}
