package token

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"karya":  FUNCTION,
	"rakha":  LET,
	"yedi":   IF,
	"natra":  ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"firta": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
