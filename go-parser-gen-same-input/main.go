package main

import (
	"fmt"

	"github.com/alexgarzao/antlr-sandbox/go-parser-gen-same-input/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var hello = `
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
`

func main() {
	goParser(hello)
}

// calc takes a string expression and returns the evaluated result.
func goParser(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewGoLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener goListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.SourceFile())
}

type goListener struct {
	*parser.BaseGoParserListener
}

func (l *goListener) VisitTerminal(node antlr.TerminalNode) {
	if node.GetText() != "<EOF>" {
		fmt.Print(node.GetText(), " ")
	}
}

func (l *goListener) EnterEos(ctx *parser.EosContext) {
	fmt.Println()
}
