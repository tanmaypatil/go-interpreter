package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	// Identifiers + Literals 
	IDENT = "IDENT" // add , foobar ,x ,y 
	INT = "INT" // 12345
	// operators 
	ASSIGN = "="
	PLUS = "+"
	// Delimiters 
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	// keywords 
	FUNCTION = "FUNCTION"
	LET = "LET"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"let": LET,
	"return" : RETURN,
} 

func LookUpIdent(ident string)TokenType {
	if tok,ok := keywords[ident];ok{
		return tok
	}
	return IDENT
}


