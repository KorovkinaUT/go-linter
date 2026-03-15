package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/KorovkinaUT/go-linter/internal/rules"
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

	rules := rules.DefaultRules()

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)

		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		if !isLogCall(pass, selector) {
			return
		}

		msg, ok := extractLogMessage(call)
		if !ok {
			return
		}

		for _, rule := range rules {
			if err := rule.Check(msg); err != "" {
				pass.Reportf(call.Pos(), "%s", err)
			}
		}
	})

	return nil, nil
}

func extractLogMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	firstArg := call.Args[0]
	// Check if first agrument is literal
	lit, ok := firstArg.(*ast.BasicLit)
	if !ok {
		return "", false
	}

	// Check if first agrument is string literal
	if lit.Kind != token.STRING {
		return "", false
	}

	return strings.Trim(lit.Value, `"`), true
}
