package token

type TokType string

type Token struct {
	Type    TokType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operaters
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Delimiters
	LPAREN    = "("
	RPAREN    = ")"
	SEMICOLON = ";"

	// Keywords
	PRINT = "PRINT"
	LET   = "LET"
)

var keywords = map[string]TokType{
	"print": PRINT,
	"let":   LET,
}

func LookUpIdent(ident string) TokType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
