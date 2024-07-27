package main

import (
	"fmt"
	"gojson/pkg/jparser"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: jparser <file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	parser := jparser.NewJParser(filePath)
	parser.Lexer.Lex()

}
