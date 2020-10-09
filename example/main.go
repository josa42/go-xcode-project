package main

import (
	"fmt"
	"io/ioutil"

	"github.com/josa42/go-xcode-project/pbxproj/formatter"
	"github.com/josa42/go-xcode-project/pbxproj/lexer"
	"github.com/josa42/go-xcode-project/pbxproj/parser"
)

func main() {
	b, _ := ioutil.ReadFile("pbxproj/parser/testdata/project.pbxproj")
	l := lexer.New(string(b))
	p := parser.New(l)

	ast, _ := p.Parse()

	fmt.Println(formatter.Format(ast))
}
