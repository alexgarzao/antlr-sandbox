package main

import (
	"fmt"

	"github.com/alexgarzao/antlr-sandbox/go-parser/parser"
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
	fmt.Println("1 + 2 * 3 =", calc(hello))
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) int {
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

	return 1
}

type goListener struct {
	*parser.BaseGoParserListener
}

func (l *goListener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Println("#### ", node.GetText())
	fmt.Print(node.GetText(), " ")
}

func (l *goListener) EnterEos(ctx *parser.EosContext) {
	// fmt.Println("//// ", ctx.GetText())
	fmt.Println()
}
