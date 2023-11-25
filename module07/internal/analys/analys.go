package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func Analys(filepath string) (*AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath,
		nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	analys := &AnalysResult{
		DeclCount:    0,
		CallCount:    0,
		AssignCount:  0,
		ImportsCount: 0,
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.GenDecl:
			analys.DeclCount++
		case *ast.CallExpr:
			analys.CallCount++
		case *ast.AssignStmt:
			analys.AssignCount++
		case *ast.ImportSpec:
			analys.ImportsCount++
		}
		return true
	})

	if analys.ImportsCount != 0 {
		analys.DeclCount--
	}

	return analys, nil
}

func main() {
	fmt.Println(Analys("module07/internal/convertor/convertor.go"))
}
