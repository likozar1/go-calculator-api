package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"github.com/go-calculator-api/pkg/calculator"
)

const Regex = "-?[0-9]+|[*+/-]"

var Expr string

func init() {
	flag.StringVar(&Expr, "ex", "", "Provide mathematical expression; example: '-10 + 15 * 3'")
}

func main() {
	flag.Parse()

	if Expr == "" {
		chexpr := make(chan string, 1)

		defer close(chexpr)

		go readStdin(chexpr)

		Expr = <-chexpr
	}

	result := calculator.CalculateParsedExpression(calculator.ParseExpression(Expr))
	fmt.Println(Expr + " = " + result)
}

func readStdin(c chan string) {
	fmt.Print("Enter expression: ")
	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic(err)
	}

	c <- string(b)
}
