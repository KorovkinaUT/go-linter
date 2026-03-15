package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func run(pass *analysis.Pass) (any, error) {
	inspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)

		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		if isLogCall(pass, selector) {
			pass.Reportf(call.Pos(), "log call detected: %s", selector.Sel.Name)
		}
	})

	return nil, nil
}
