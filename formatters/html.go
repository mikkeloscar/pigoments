package formatters

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/mikkeloscar/pigoments/lexers"
)

var cssClasses = map[lexers.TokenType]string{
	lexers.TokenError: "err",
	// comments
	lexers.TokenComment:          "c",
	lexers.TokenCommentMultiline: "cm",
	lexers.TokenCommentPreproc:   "cp",
	lexers.TokenCommentSingle:    "c1",
	lexers.TokenCommentSpecial:   "cs",
	// generics
	lexers.TokenGenericDeleted:    "gd",
	lexers.TokenGenericEmph:       "ge",
	lexers.TokenGenericError:      "gr",
	lexers.TokenGenericHeading:    "gh",
	lexers.TokenGenericInserted:   "gi",
	lexers.TokenGenericOutput:     "go",
	lexers.TokenGenericPrompt:     "gp",
	lexers.TokenGenericStrong:     "gs",
	lexers.TokenGenericSubheading: "gu",
	lexers.TokenGenericTraceback:  "gt",
	// keyword
	lexers.TokenKeyword:            "k",
	lexers.TokenKeywordConstant:    "kc",
	lexers.TokenKeywordDeclaration: "kd",
	lexers.TokenKeywordNamespace:   "kn",
	lexers.TokenKeywordPseudo:      "kp",
	lexers.TokenKeywordReserved:    "kr",
	lexers.TokenKeywordType:        "kt",
	// literal
	lexers.TokenLiteralNumber:            "m",
	lexers.TokenLiteralString:            "s",
	lexers.TokenLiteralNumberFloat:       "mf",
	lexers.TokenLiteralNumberHex:         "mh",
	lexers.TokenLiteralNumberInteger:     "mi",
	lexers.TokenLiteralNumberOct:         "mo",
	lexers.TokenLiteralStringBacktick:    "sb",
	lexers.TokenLiteralStringChar:        "sc",
	lexers.TokenLiteralStringDoc:         "sd",
	lexers.TokenLiteralStringDouble:      "s2",
	lexers.TokenLiteralStringEscape:      "se",
	lexers.TokenLiteralStringHeredoc:     "sh",
	lexers.TokenLiteralStringInterpol:    "si",
	lexers.TokenLiteralStringOther:       "sx",
	lexers.TokenLiteralStringRegex:       "sr",
	lexers.TokenLiteralStringSingle:      "s1",
	lexers.TokenLiteralStringSymbol:      "ss",
	lexers.TokenLiteralNumberIntegerLong: "il",
	// name
	lexers.TokenNameAttribute:        "na",
	lexers.TokenNameBuiltin:          "nb",
	lexers.TokenNameClass:            "nc",
	lexers.TokenNameConstant:         "no",
	lexers.TokenNameDecorator:        "nd",
	lexers.TokenNameEntity:           "ni",
	lexers.TokenNameException:        "ne",
	lexers.TokenNameFunction:         "nf",
	lexers.TokenNameLabel:            "nl",
	lexers.TokenNameNamespace:        "nn",
	lexers.TokenNameTag:              "nt",
	lexers.TokenName:                 "nx",
	lexers.TokenNameVariable:         "nv",
	lexers.TokenNameBuiltinPseudo:    "bp",
	lexers.TokenNameVariableClass:    "vc",
	lexers.TokenNameVariableGlobal:   "vg",
	lexers.TokenNameVariableInstance: "vi",
	// operator
	lexers.TokenOperator:     "o",
	lexers.TokenOperatorWord: "ow",
	// text
	lexers.TokenTextWhitespace: "w",
	// punctuation
	lexers.TokenPunctuation: "p",
}

var escapeHTMLTable = map[rune]string{
	'&':  "&amp;",
	'<':  "&lt;",
	'>':  "&gt;",
	'"':  "&quot;",
	'\'': "&#39;",
}

const (
	wrapper      = `<div class="%s"><pre>%s</pre></div>`
	linenosTable = `<table class="highlighttable"><tr><td class="linenos">
<div class="linenodeiv"><pre>%s</pre></pre></div></td><td class="code">%s</td></tr></table>`
)

// HTMLFormatter can create html formatted output
type HTMLFormatter struct {
	CSSClass string
	LineNos  bool
}

// Generate can generate the html
func (f *HTMLFormatter) Generate(l *lexers.Lexer) string {
	// set default if value not chosen by caller
	f.defaults()

	var buf bytes.Buffer
	var token *lexers.Token

	linenos := 0

	for {
		token = l.NextToken()
		fmt.Printf("%#v\n", token)
		if token == nil {

			code := buf.String()

			// handle input ending on a newline like a real editor
			if f.LineNos && code[len(code)-1] != '\n' {
				linenos++
			}

			return f.format(fmt.Sprintf(wrapper, f.CSSClass, code), linenos)
		}

		if f.LineNos {
			for _, c := range token.Value {
				if c == '\n' {
					linenos++
				}
			}
		}

		if css, ok := cssClasses[token.Type]; ok {
			buf.WriteString(fmt.Sprintf(`<span class="%s">`, css))
			buf.WriteString(escape(token.Value))
			buf.WriteString("</span>")
		} else {
			buf.WriteString(token.Value)
		}
	}
}

// initialize default values if needed
func (f *HTMLFormatter) defaults() {
	if f.CSSClass == "" {
		f.CSSClass = "highlight"
	}
}

// do the final formatting depending on linenumbers should be included or not
func (f *HTMLFormatter) format(code string, linenos int) string {
	if !f.LineNos {
		return code
	}

	var buf bytes.Buffer
	var diff int
	width := numDigits(linenos)
	for i := 1; i <= linenos; i++ {
		diff = width - numDigits(i)
		for j := 1; j <= diff; j++ {
			buf.WriteRune(' ')
		}
		buf.WriteString(strconv.Itoa(i))
		if i < linenos {
			buf.WriteRune('\n')
		}
	}

	return fmt.Sprintf(linenosTable, buf.String(), code)
}

// count the number of digits in an integer
func numDigits(n int) int {
	switch {
	case n <= 0:
		return 0
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 1000:
		return 3
	case n < 10000:
		return 4
	case n < 100000:
		return 5
	default:
		return 6
	}
}

func escape(str string) string {
	var buf bytes.Buffer

	for _, c := range str {
		if e, ok := escapeHTMLTable[c]; ok {
			buf.WriteString(e)
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
