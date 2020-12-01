package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestMyListener1(t *testing.T) {
	is := antlr.NewInputStream(`
	int summation(int a, int b) {
         int sum;
         sum = a+b;
         return sum;
	}

	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	listener := NewMyListener()
	p.AddParseListener(listener)
	tree := p.CompilationUnit()
	fmt.Println(tree.ToStringTree([]string{}, p))
}

/*
func TestMyListener2(t *testing.T) {
	is := antlr.NewInputStream(`
	#include <stdio.h>
	main() {
	    printf("Hello world");
	}
	`)
	lexer := NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewCParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(NewMyListener(), p.Declaration())
}

func TestMyListener3(t *testing.T) {
	fmt.Println("TestMyListener3")
}
*/
