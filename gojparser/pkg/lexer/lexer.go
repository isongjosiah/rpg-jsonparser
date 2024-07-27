package lexer

import (
	"fmt"
	"github.com/pkg/errors"
	"gojson/pkg"
	"os"
	"strings"
)

type Lexer struct {
	filePath string
	Tokens   []pkg.Token
}

func NewLexer(filePath string) *Lexer {
	return &Lexer{
		filePath: filePath,
		Tokens:   make([]pkg.Token, 0),
	}
}

func (l *Lexer) Lex() error {

	// read the file content
	data, err := os.ReadFile(l.filePath)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}

	return l.tokenize(data)

}

func (l *Lexer) tokenize(data []byte) error {

	dataString := string(data)
	dataString = strings.TrimSpace(dataString)
	dss := strings.Split(dataString, "\n")

	for _, ds := range dss {
		if len(ds) == 0 { // TODO: revisit
			continue
		}
		ds = strings.Trim(ds, "\n")
		fmt.Println(ds)
		switch ds {
		case pkg.L_BRACE, pkg.R_BRACE, pkg.COLON, pkg.L_BRACKET, pkg.R_BRACKET, pkg.COMMA:
			l.Tokens = append(l.Tokens, pkg.Token(ds))
		default:
			// either be a key or value
			switch {
			case l.lastToken() == pkg.COMMA || l.lastToken() == pkg.R_BRACKET && (strings.Contains(ds, `"`)): // it is an identifier
				l.Tokens = append(l.Tokens, pkg.IDENTIFIER)

			case l.lastToken() == pkg.COLON: // it is a value
				l.Tokens = append(l.Tokens, pkg.LITERAL)
			}
		}
	}

	fmt.Println(l.Tokens)
	return nil
}

func (l *Lexer) lastToken() pkg.Token {
	if len(l.Tokens) > 0 {
		return l.Tokens[len(l.Tokens)-1]
	}
	return pkg.INVALID
}

/*
{
	"key": "value",
	"key2": "x"
}
*/
