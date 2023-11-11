package token

type TokenType string

type Token struct {
	Type      TokenType
	Literal   string
	RowNumber int
	ColNumber int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT    = "IDENT"
	NUMBER   = "NUMBER"
	STRING   = "STRING"
	OPERATOR = "OPERATOR"

	UNDERSCORE = "_"
	QUESTION   = "?"

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	PERCENT   = "%"
	AMPERSAND = "&"
	PIPE      = "|"
	LT        = "<"
	GT        = ">"

	DOT       = "."
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	AT        = "@"
	DOLLAR    = "$"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	NIL    = "NIL"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	ELSEIF = "ELSEIF"
	UNLESS = "UNLESS"
	RETURN = "RETURN"
	DO     = "DO"
	END    = "END"
	WHILE  = "WHILE"
)

var keywords = map[string]TokenType{
	"nil":    NIL,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"elseif": ELSEIF,
	"unless": UNLESS,
	"return": RETURN,
	"do":     DO,
	"end":    END,
	"while":  WHILE,
}

var operator_symbols = map[string]TokenType{
	"<": LT,
	">": GT,
	"=": ASSIGN,
	"+": PLUS,
	"-": MINUS,
	"!": BANG,
	"*": ASTERISK,
	"/": SLASH,
	"&": AMPERSAND,
	"%": PERCENT,
	"|": PIPE,
}

var symbols = map[string]TokenType{
	".": DOT,
	",": COMMA,
	":": COLON,
	";": SEMICOLON,
	"@": AT,
	"$": DOLLAR,
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

func IsOperatorSymbol(ch string) bool {
	if _, ok := operator_symbols[ch]; ok {
		return true
	}

	return false
}

func IsSymbol(ch string) bool {
	if _, ok := symbols[ch]; ok {
		return true
	}

	return false
}

func LookupSymbol(ch string) TokenType {
	if tok, ok := symbols[ch]; ok {
		return tok
	}

	return ILLEGAL
}
