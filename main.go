package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"monkey/ast"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)

	flag.Parse()
	filePath := flag.Args()[0]

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failure to read file '%s'. Err: %s", string(contents), err)
	}
	l := lexer.New(string(contents))
	p := parser.New(l)
	program := p.ParseProgram()
	evaluateAst(program)
}

// Evaluate the AST with evaluator
func evaluateAst(program *ast.Program) object.Object {
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}
