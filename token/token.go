package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"def":    DEF,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
	"while":  WHILE,
	"break":  BREAK,
	"class":  CLASS,
	"self":   SELF,
	"none":   NONE,
	"pass":   PASS,
	"and":    AND,
	"or":     OR,
	"not":    NOT,
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

	// Identifiers
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	MOD    = "%"
	EXP    = "**"

	// Comparison Operators
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = "=>"

	// Delimiters
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	COMMA     = ","
	COLON     = ":"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	IF     = "if"
	ELSE   = "else"
	DEF    = "def"
	RETURN = "return"
	TRUE   = "true"
	FALSE  = "false"
	FOR    = "for"
	WHILE  = "while"
	BREAK  = "break"
	CLASS  = "class"
	SELF   = "self"
	AND    = "and"
	OR     = "or"
	NOT    = "not"
	NONE   = "none"
	PASS   = "pass"
)
