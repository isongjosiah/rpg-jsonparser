package jparser

import "gojson/pkg/lexer"

type JParser struct {
	Lexer *lexer.Lexer
}

func NewJParser(filePath string) *JParser {
	return &JParser{
		Lexer: lexer.NewLexer(filePath),
	}
}
