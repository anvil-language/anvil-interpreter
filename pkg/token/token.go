package token

type TokType string

type Token struct {
	Type    TokType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NOT_EQ  = "!="
	EQ      = "=="

	// Identifiers and literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operaters
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	SEMICOLON = ";"

	// conditionals
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Keywords
	PRINT = "PRINT"
	LET   = "LET"
	IF    = "IF"
	ELSE  = "ELSE"
)

var keywords = map[string]TokType{
	"print": PRINT,
	"true":  TRUE,
	"false": FALSE,
	"let":   LET,
	"if":    IF,
	"else":  ELSE,
}

func LookUpIdent(ident string) TokType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
