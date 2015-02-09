package lexers

// Token describes a token with the token type and value
type Token struct {
	Type  TokenType
	Value string
}

// TokenType defines a type of token
type TokenType int

const (
	// TokenError defines an error token
	TokenError TokenType = iota

	// comments

	// TokenComment defines a general comment token
	TokenComment
	// TokenCommentMultiline defines a multiline comment token
	TokenCommentMultiline
	// TokenCommentPreproc defines a preproc comment token
	TokenCommentPreproc
	// TokenCommentSingle defines a single line comment token
	TokenCommentSingle
	// TokenCommentSpecial defines a special comment token
	TokenCommentSpecial

	// generics

	// TokenGenericDeleted defines a generic deleted token
	TokenGenericDeleted
	// TokenGenericEmph defines a generic emphasis token
	TokenGenericEmph
	// TokenGenericError defines a generic error token
	TokenGenericError
	// TokenGenericHeading defines a generic heading token
	TokenGenericHeading
	// TokenGenericInserted defines a generic inserted token
	TokenGenericInserted
	// TokenGenericOutput defines a generic output token
	TokenGenericOutput
	// TokenGenericPrompt defines a generic prompt token
	TokenGenericPrompt
	// TokenGenericStrong defines a generic strong type token
	TokenGenericStrong
	// TokenGenericSubheading defines a generic subheading token
	TokenGenericSubheading
	// TokenGenericTraceback defines a generic traceback token
	TokenGenericTraceback

	// keyword

	// TokenKeyword defines a keyword token
	TokenKeyword
	// TokenKeywordConstant defines a constant keyword token
	TokenKeywordConstant
	// TokenKeywordDeclaration defines a keyword declaration token
	TokenKeywordDeclaration
	// TokenKeywordNamespace defines a namespace keyword token
	TokenKeywordNamespace
	// TokenKeywordPseudo defines a pseudo keyword token
	TokenKeywordPseudo
	// TokenKeywordReserved defines a reserved keyword token
	TokenKeywordReserved
	// TokenKeywordType defines a reserved type keyword token
	TokenKeywordType

	// literal

	// TokenLiteralNumber defines a literal number token
	TokenLiteralNumber
	// TokenLiteralString defines a literal string token
	TokenLiteralString
	// TokenLiteralNumberFloat defines a literal float number token
	TokenLiteralNumberFloat
	// TokenLiteralNumberHex defines a literal heximal number token
	TokenLiteralNumberHex
	// TokenLiteralNumberInteger defines a literal integer number token
	TokenLiteralNumberInteger
	// TokenLiteralNumberOct defines a literal octal number token
	TokenLiteralNumberOct
	// TokenLiteralNumberIntegerLong defines a literal long integer number token
	TokenLiteralNumberIntegerLong
	// TokenLiteralStringBacktick defines a literal string backtick token
	TokenLiteralStringBacktick
	// TokenLiteralStringChar defines a literal string char token
	TokenLiteralStringChar
	// TokenLiteralStringDoc defines a literal doc string token
	TokenLiteralStringDoc
	// TokenLiteralStringDouble defines a literal double string token
	TokenLiteralStringDouble
	// TokenLiteralStringEscape defines a literal string escape token
	TokenLiteralStringEscape
	// TokenLiteralStringHeredoc defines a literal Heredoc string token
	TokenLiteralStringHeredoc
	// TokenLiteralStringInterpol defines a literal interpol string token
	TokenLiteralStringInterpol
	// TokenLiteralStringOther defines a literal string other token
	TokenLiteralStringOther
	// TokenLiteralStringRegex defines a literal regex string token
	TokenLiteralStringRegex
	// TokenLiteralStringSingle defines a literal single string token
	TokenLiteralStringSingle
	// TokenLiteralStringSymbol defines a literal string symbol token
	TokenLiteralStringSymbol

	// name

	// TokenName defines a name token
	TokenName
	// TokenNameAttribute defines a name attribute token
	TokenNameAttribute
	// TokenNameBuiltin defines a builtin name token
	TokenNameBuiltin
	// TokenNameClass defines a class name token
	TokenNameClass
	// TokenNameConstant defines a constant name token
	TokenNameConstant
	// TokenNameDecorator defines a decorator name token
	TokenNameDecorator
	// TokenNameEntity defines an entity name token
	TokenNameEntity
	// TokenNameException defines an exception name token
	TokenNameException
	// TokenNameFunction defines a function name token
	TokenNameFunction
	// TokenNameLabel defines a label name token
	TokenNameLabel
	// TokenNameNamespace defines a namespace name token
	TokenNameNamespace
	// TokenNameTag defines a tag name token
	TokenNameTag
	// TokenNameVariable defines a variable name token
	TokenNameVariable
	// TokenNameBuiltinPseudo defines a builtin pseudo name token
	TokenNameBuiltinPseudo
	// TokenNameVariableClass defines a variable class name token
	TokenNameVariableClass
	// TokenNameVariableGlobal defines a variable global name token
	TokenNameVariableGlobal
	// TokenNameVariableInstance defines a variable instance name token
	TokenNameVariableInstance

	// operator

	// TokenOperator defines an operator token
	TokenOperator
	// TokenOperatorWord defines an operator word token
	TokenOperatorWord

	// text

	// TokenText defines a text token, this is the default token for
	// non-highlighted text
	TokenText
	// TokenTextWhitespace defines a whitespace text token
	TokenTextWhitespace

	// TokenPunctuation defines a puctuation token
	TokenPunctuation
)
