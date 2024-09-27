package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def": DEF,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"
	STRING  = "STRING"
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	MUL     = "*"
	DIV     = "/"
	LPAREN  = "("
	RPAREN  = ")"
	COMMA   = ","
	IF      = "if"
	ELSE    = "else"
	DEF     = "def"
)
