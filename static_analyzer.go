package static_analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "static_analyzer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "static_analyzer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BlockStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BlockStmt:
			check(pass, n)
		}
	})

	return nil, nil
}

func check(pass *analysis.Pass, n *ast.BlockStmt) {
	sz := len(n.List)
	uninitializedVariables := make([][]string, sz)
	assignedVariables := make([][]string, sz)
	for i := 0; i < sz; i++ {
		uninitializedVariables[0] = make([]string, 0)
		assignedVariables[0] = make([]string, 0)
	}

	for i, node := range n.List {
		switch node := node.(type) {
		case *ast.DeclStmt:
			isUninitialized, names, err := checkDecl(node.Decl)
			if err == nil && isUninitialized {
				for _, name := range names {
					uninitializedVariables[i] = append(uninitializedVariables[i], name.Name)
				}
			}
		}
	}

	for i, vars := range uninitializedVariables {
		if len(vars) == 0 {
			continue
		}

		checkAssign(pass, n.List[i+1], vars)
	}
}

func checkDecl(n ast.Decl) (bool, []*ast.Ident, error) {
	ast.Print(token.NewFileSet(), n)

	genDecl := n.(*ast.GenDecl)
	if len(genDecl.Specs[0].(*ast.ValueSpec).Values) != 0 {
		return false, nil, nil
	}

	names := genDecl.Specs[0].(*ast.ValueSpec).Names

	return true, names, nil
}

func checkAssign(pass *analysis.Pass, node ast.Stmt, vars []string) {
	switch node := node.(type) {
	case *ast.AssignStmt:
		pass.Reportf(node.Pos(), "not implemented")
	case *ast.IfStmt:
		pass.Reportf(node.Pos(), "not implemented")
	default:
		pass.Reportf(node.Pos(), "variable is not assigned")
	}
}










