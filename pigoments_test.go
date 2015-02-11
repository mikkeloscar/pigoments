package pigoments

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/mikkeloscar/pigoments/lexers"
)

const testDir = "tests"

func expectedTokens(tokensFile string, t *testing.T) []string {
	// read list of expected tokens
	content, err := ioutil.ReadFile(path.Join(testDir, tokensFile))
	if err != nil {
		t.Error(err)
	}
	expectedTokens := strings.Split(string(content), "\n")
	expectedTokens = expectedTokens[0 : len(expectedTokens)-1]

	return expectedTokens
}

func TestLexers(t *testing.T) {
	var err error
	fis, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Error(err)
	}

	lexerTests := make(map[string]string, len(fis))

	for i := 0; i < len(fis)-1; i += 2 {
		str1 := strings.Split(fis[i].Name(), ".")
		if str1[len(str1)-1] == "token" {
			lexerTests[fis[i+1].Name()] = fis[i].Name()
		} else {
			lexerTests[fis[i].Name()] = fis[i+1].Name()
		}
	}

	var lang string
	var l *lexers.Lexer
	var codeContent []byte
	var expTokens []string
	var token *lexers.Token
	var i int
	for code, tokens := range lexerTests {
		lang = ""
		// get lang from file name
		for _, c := range code {
			if c == '.' || c == '_' {
				break
			}
			lang += string(c)
		}

		expTokens = expectedTokens(tokens, t)

		// read code to be lexed
		codeContent, err = ioutil.ReadFile(path.Join(testDir, code))
		if err != nil {
			t.Error(err)
		}

		// spawn lexer
		l = lexers.Lex(string(codeContent), lang)

		i = 0
		for {
			token = l.NextToken()
			if token == nil {
				break
			}

			if i == len(expTokens) {
				t.Errorf("code generated more tokens than expected")
				return
			}

			if expected, ok := tokensMap[expTokens[i]]; ok {
				if expected != token.Type {
					t.Errorf("invalid token: %s, expected: %s", tokToStr[token.Type], tokToStr[expected])
				}
			} else {
				t.Errorf("invalid token: '%s' not found", expTokens[i])
			}

			i++
		}
	}
}

var tokToStr = map[lexers.TokenType]string{
	lexers.TokenError:                    "TokenError",
	lexers.TokenComment:                  "TokenComment",
	lexers.TokenCommentMultiline:         "TokenCommentMultiline",
	lexers.TokenCommentPreproc:           "TokenCommentPreproc",
	lexers.TokenCommentSingle:            "TokenCommentSingle",
	lexers.TokenCommentSpecial:           "TokenCommentSpecial",
	lexers.TokenGenericDeleted:           "TokenGenericDeleted",
	lexers.TokenGenericEmph:              "TokenGenericEmph",
	lexers.TokenGenericError:             "TokenGenericError",
	lexers.TokenGenericHeading:           "TokenGenericHeading",
	lexers.TokenGenericInserted:          "TokenGenericInserted",
	lexers.TokenGenericOutput:            "TokenGenericOutput",
	lexers.TokenGenericPrompt:            "TokenGenericPrompt",
	lexers.TokenGenericStrong:            "TokenGenericStrong",
	lexers.TokenGenericSubheading:        "TokenGenericSubheading",
	lexers.TokenGenericTraceback:         "TokenGenericTraceback",
	lexers.TokenKeyword:                  "TokenKeyword",
	lexers.TokenKeywordConstant:          "TokenKeywordConstant",
	lexers.TokenKeywordDeclaration:       "TokenKeywordDeclaration",
	lexers.TokenKeywordNamespace:         "TokenKeywordNamespace",
	lexers.TokenKeywordPseudo:            "TokenKeywordPseudo",
	lexers.TokenKeywordReserved:          "TokenKeywordReserved",
	lexers.TokenKeywordType:              "TokenKeywordType",
	lexers.TokenLiteralNumber:            "TokenLiteralNumber",
	lexers.TokenLiteralString:            "TokenLiteralString",
	lexers.TokenLiteralNumberFloat:       "TokenLiteralNumberFloat",
	lexers.TokenLiteralNumberHex:         "TokenLiteralNumberHex",
	lexers.TokenLiteralNumberInteger:     "TokenLiteralNumberInteger",
	lexers.TokenLiteralNumberOct:         "TokenLiteralNumberOct",
	lexers.TokenLiteralNumberIntegerLong: "TokenLiteralNumberIntegerLong",
	lexers.TokenLiteralStringBacktick:    "TokenLiteralStringBacktick",
	lexers.TokenLiteralStringChar:        "TokenLiteralStringChar",
	lexers.TokenLiteralStringDoc:         "TokenLiteralStringDoc",
	lexers.TokenLiteralStringDouble:      "TokenLiteralStringDouble",
	lexers.TokenLiteralStringEscape:      "TokenLiteralStringEscape",
	lexers.TokenLiteralStringHeredoc:     "TokenLiteralStringHeredoc",
	lexers.TokenLiteralStringInterpol:    "TokenLiteralStringInterpol",
	lexers.TokenLiteralStringOther:       "TokenLiteralStringOther",
	lexers.TokenLiteralStringRegex:       "TokenLiteralStringRegex",
	lexers.TokenLiteralStringSingle:      "TokenLiteralStringSingle",
	lexers.TokenLiteralStringSymbol:      "TokenLiteralStringSymbol",
	lexers.TokenName:                     "TokenName",
	lexers.TokenNameAttribute:            "TokenNameAttribute",
	lexers.TokenNameBuiltin:              "TokenNameBuiltin",
	lexers.TokenNameClass:                "TokenNameClass",
	lexers.TokenNameConstant:             "TokenNameConstant",
	lexers.TokenNameDecorator:            "TokenNameDecorator",
	lexers.TokenNameEntity:               "TokenNameEntity",
	lexers.TokenNameException:            "TokenNameException",
	lexers.TokenNameFunction:             "TokenNameFunction",
	lexers.TokenNameLabel:                "TokenNameLabel",
	lexers.TokenNameNamespace:            "TokenNameNamespace",
	lexers.TokenNameTag:                  "TokenNameTag",
	lexers.TokenNameVariable:             "TokenNameVariable",
	lexers.TokenNameBuiltinPseudo:        "TokenNameBuiltinPseudo",
	lexers.TokenNameVariableClass:        "TokenNameVariableClass",
	lexers.TokenNameVariableGlobal:       "TokenNameVariableGlobal",
	lexers.TokenNameVariableInstance:     "TokenNameVariableInstance",
	lexers.TokenOperator:                 "TokenOperator",
	lexers.TokenOperatorWord:             "TokenOperatorWord",
	lexers.TokenText:                     "TokenText",
	lexers.TokenTextWhitespace:           "TokenTextWhitespace",
	lexers.TokenPunctuation:              "TokenPunctuation",
}

var tokensMap = map[string]lexers.TokenType{
	"TokenError":                    lexers.TokenError,
	"TokenComment":                  lexers.TokenComment,
	"TokenCommentMultiline":         lexers.TokenCommentMultiline,
	"TokenCommentPreproc":           lexers.TokenCommentPreproc,
	"TokenCommentSingle":            lexers.TokenCommentSingle,
	"TokenCommentSpecial":           lexers.TokenCommentSpecial,
	"TokenGenericDeleted":           lexers.TokenGenericDeleted,
	"TokenGenericEmph":              lexers.TokenGenericEmph,
	"TokenGenericError":             lexers.TokenGenericError,
	"TokenGenericHeading":           lexers.TokenGenericHeading,
	"TokenGenericInserted":          lexers.TokenGenericInserted,
	"TokenGenericOutput":            lexers.TokenGenericOutput,
	"TokenGenericPrompt":            lexers.TokenGenericPrompt,
	"TokenGenericStrong":            lexers.TokenGenericStrong,
	"TokenGenericSubheading":        lexers.TokenGenericSubheading,
	"TokenGenericTraceback":         lexers.TokenGenericTraceback,
	"TokenKeyword":                  lexers.TokenKeyword,
	"TokenKeywordConstant":          lexers.TokenKeywordConstant,
	"TokenKeywordDeclaration":       lexers.TokenKeywordDeclaration,
	"TokenKeywordNamespace":         lexers.TokenKeywordNamespace,
	"TokenKeywordPseudo":            lexers.TokenKeywordPseudo,
	"TokenKeywordReserved":          lexers.TokenKeywordReserved,
	"TokenKeywordType":              lexers.TokenKeywordType,
	"TokenLiteralNumber":            lexers.TokenLiteralNumber,
	"TokenLiteralString":            lexers.TokenLiteralString,
	"TokenLiteralNumberFloat":       lexers.TokenLiteralNumberFloat,
	"TokenLiteralNumberHex":         lexers.TokenLiteralNumberHex,
	"TokenLiteralNumberInteger":     lexers.TokenLiteralNumberInteger,
	"TokenLiteralNumberOct":         lexers.TokenLiteralNumberOct,
	"TokenLiteralNumberIntegerLong": lexers.TokenLiteralNumberIntegerLong,
	"TokenLiteralStringBacktick":    lexers.TokenLiteralStringBacktick,
	"TokenLiteralStringChar":        lexers.TokenLiteralStringChar,
	"TokenLiteralStringDoc":         lexers.TokenLiteralStringDoc,
	"TokenLiteralStringDouble":      lexers.TokenLiteralStringDouble,
	"TokenLiteralStringEscape":      lexers.TokenLiteralStringEscape,
	"TokenLiteralStringHeredoc":     lexers.TokenLiteralStringHeredoc,
	"TokenLiteralStringInterpol":    lexers.TokenLiteralStringInterpol,
	"TokenLiteralStringOther":       lexers.TokenLiteralStringOther,
	"TokenLiteralStringRegex":       lexers.TokenLiteralStringRegex,
	"TokenLiteralStringSingle":      lexers.TokenLiteralStringSingle,
	"TokenLiteralStringSymbol":      lexers.TokenLiteralStringSymbol,
	"TokenName":                     lexers.TokenName,
	"TokenNameAttribute":            lexers.TokenNameAttribute,
	"TokenNameBuiltin":              lexers.TokenNameBuiltin,
	"TokenNameClass":                lexers.TokenNameClass,
	"TokenNameConstant":             lexers.TokenNameConstant,
	"TokenNameDecorator":            lexers.TokenNameDecorator,
	"TokenNameEntity":               lexers.TokenNameEntity,
	"TokenNameException":            lexers.TokenNameException,
	"TokenNameFunction":             lexers.TokenNameFunction,
	"TokenNameLabel":                lexers.TokenNameLabel,
	"TokenNameNamespace":            lexers.TokenNameNamespace,
	"TokenNameTag":                  lexers.TokenNameTag,
	"TokenNameVariable":             lexers.TokenNameVariable,
	"TokenNameBuiltinPseudo":        lexers.TokenNameBuiltinPseudo,
	"TokenNameVariableClass":        lexers.TokenNameVariableClass,
	"TokenNameVariableGlobal":       lexers.TokenNameVariableGlobal,
	"TokenNameVariableInstance":     lexers.TokenNameVariableInstance,
	"TokenOperator":                 lexers.TokenOperator,
	"TokenOperatorWord":             lexers.TokenOperatorWord,
	"TokenText":                     lexers.TokenText,
	"TokenTextWhitespace":           lexers.TokenTextWhitespace,
	"TokenPunctuation":              lexers.TokenPunctuation,
}
