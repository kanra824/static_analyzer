package static_analyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
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
	// グローバルの変数は一旦無視して、関数ブロック中の変数について検出
	decls := pass.Files[0].Decls
	for _, decl := range decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			variables := make(map[string]bool)
			block := decl.Body
			checkBlockStmt(pass, block, variables)
		}
	}

	return nil, nil
}


func checkStmt(pass *analysis.Pass, node ast.Stmt, variables map[string]bool) {
	switch node := node.(type) {
	case *ast.BlockStmt:
		newVariables := make(map[string]bool)
		for k, v := range variables {
			if v {
				newVariables[k] = v
			}
		}
		checkBlockStmt(pass, node, newVariables)
		// TODO: ブロックの外側で定義された変数にブロックの中で代入している場合を検出
	case *ast.DeclStmt:
		checkDeclStmt(node.Decl, variables)
	case *ast.AssignStmt:
		checkAssignStmt(pass, node, variables)
	case *ast.IfStmt:
		if node.Init != nil {
			checkStmt(pass, node.Init, variables)
		}
		checkExpr(pass, node.Cond, variables)
		checkStmt(pass, node.Body, variables)
	case *ast.ForStmt:
		if node.Init != nil {
			checkStmt(pass, node.Init, variables)
		}
		if node.Cond != nil {
			checkExpr(pass, node.Cond, variables)
		}
		if node.Post != nil {
			checkStmt(pass, node.Post, variables)
		}
		checkStmt(pass, node.Body, variables)
	case *ast.ReturnStmt:
		for _, ch := range node.Results {
			checkExpr(pass, ch, variables)
		}
	case *ast.IncDecStmt:
		checkExpr(pass, node.X, variables)
	case *ast.SwitchStmt:
		if node.Init != nil {
			checkStmt(pass, node.Init, variables)
		}
		if node.Tag != nil {
			checkExpr(pass, node.Tag, variables)
		}
		checkStmt(pass, node.Body, variables)
	case *ast.CaseClause:
		for _, exp := range node.List {
			checkExpr(pass, exp, variables)
		}
		for _, stmt := range node.Body {
			checkStmt(pass, stmt, variables)
		}
	case *ast.ExprStmt:
		checkExpr(pass, node.X, variables)
	default:
		pass.Reportf(node.Pos(), "not yet implemented")
	}
}

// ブロックごとに、文を順番にチェック
func checkBlockStmt(pass *analysis.Pass, block *ast.BlockStmt, variables map[string]bool) {
	for _, node := range block.List {
		checkStmt(pass, node, variables)
	}
}

// 宣言のなかで初期化されていない変数があるかどうかチェック
func checkDeclStmt(n ast.Decl, variables map[string]bool) {
	genDecl := n.(*ast.GenDecl)
	if len(genDecl.Specs[0].(*ast.ValueSpec).Values) != 0 {
		return
	}

	names := genDecl.Specs[0].(*ast.ValueSpec).Names
	for _, name := range names {
		variables[name.Name] = true
	}
}

func checkAssignStmt(pass *analysis.Pass, node *ast.AssignStmt, variables map[string]bool) {
	for _, expr := range node.Rhs {
		checkExpr(pass, expr, variables)
	}
	// とりあえず左辺値が単なる変数の時を考える(a[0] = 1などが未対応)
	for _, lhs := range node.Lhs {
		switch lhs := lhs.(type) {
		case *ast.Ident:
			if _, ok := variables[lhs.Name]; ok {
				variables[lhs.Name] = false
			}
		default:
			pass.Reportf(lhs.Pos(), "not yet implemented")
		}
	}
}

func checkExpr(pass *analysis.Pass, node ast.Expr, variables map[string]bool) {
	switch node := node.(type) {
	case *ast.Ident:
		if _, ok := variables[node.Name]; ok && variables[node.Name] {
			pass.Reportf(node.Pos(), "variable is not initialized")
		}
	case *ast.BinaryExpr:
		checkExpr(pass, node.X, variables)
		checkExpr(pass, node.Y, variables)
	case *ast.CallExpr:
		for _, arg := range node.Args {
			checkExpr(pass, arg, variables)
		}
	case *ast.BasicLit:
	case *ast.ParenExpr:
		checkExpr(pass, node.X, variables)
	case *ast.SelectorExpr:
		// TODO: node.Selがゼロ値かどうかを調べる
		checkExpr(pass, node.X, variables)
	default:
		pass.Reportf(node.Pos(), "not yet implemented")
	}
}
