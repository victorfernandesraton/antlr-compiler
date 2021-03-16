package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	parser "github.com/victorfernandesraton/antlr-compiler/Expr"
)

func main() {
	// gera um conjunto de input
	is := antlr.NewInputStream("(3 + 4)\n 5 + 3")

	// criar analisador lexico
	lexer := parser.NewExprLexer(is)

	// Leitura de resutados da arvore
	for {
		// pega o proximo termo disponivel
		t := lexer.NextToken()
		// se o proximo termo for o encerrador de sintaxe encerra o loop
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("token: %s\tvalue: %q\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
