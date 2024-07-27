package pkg

// Token defines the lexical dictionary for our json parser
type Token string

const (
	L_BRACE    = "{"
	R_BRACE    = "}"
	COLON      = ":"
	L_BRACKET  = "["
	R_BRACKET  = "]"
	COMMA      = ","
	IDENTIFIER = "IDENTIFIER"
	LITERAL    = "LITERAL"
	INVALID    = "INVALID"
)
