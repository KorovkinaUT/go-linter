package analyzer

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func isLogCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	methodName := selector.Sel.Name
	if !isLogMethod(methodName) {
		return false
	}

	return isStdLogPackageCall(pass, selector) ||
		isSlogPackageCall(pass, selector) ||
		isSlogLoggerMethodCall(pass, selector) ||
		isZapLoggerMethodCall(pass, selector)
}

func isStdLogPackageCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	// Get calling object identifier
	ident, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[ident]
	// Check if it is package
	pkgName, ok := obj.(*types.PkgName)
	if !ok {
		return false
	}

	// Check if it is "log" package
	return pkgName.Imported().Path() == "log"
}

func isSlogPackageCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	// Get calling object identifier
	ident, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[ident]
	// Check if it is package
	pkgName, ok := obj.(*types.PkgName)
	if !ok {
		return false
	}

	// Check if it is "slog" package
	return pkgName.Imported().Path() == "log/slog"
}

func isSlogLoggerMethodCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	typeStr := typeString(pass, selector.X)

	// Check slog.Logger type
	return typeStr == "*log/slog.Logger" || typeStr == "log/slog.Logger"
}

func isZapLoggerMethodCall(pass *analysis.Pass, selector *ast.SelectorExpr) bool {
	typeStr := typeString(pass, selector.X)

	// Checks zap.Logger type
	return typeStr == "*go.uber.org/zap.Logger" ||
		typeStr == "go.uber.org/zap.Logger" ||
		typeStr == "*go.uber.org/zap.SugaredLogger" ||
		typeStr == "go.uber.org/zap.SugaredLogger"
}

// Returns type is string with full package path
func typeString(pass *analysis.Pass, expr ast.Expr) string {
	typ := pass.TypesInfo.TypeOf(expr)
	if typ == nil {
		return ""
	}

	return types.TypeString(typ, func(pkg *types.Package) string {
		if pkg == nil {
			return ""
		}
		return pkg.Path()
	})
}

func isLogMethod(name string) bool {
	switch {
	case strings.HasPrefix(name, "Debug"),
		strings.HasPrefix(name, "Info"),
		strings.HasPrefix(name, "Warn"),
		strings.HasPrefix(name, "Error"),
		strings.HasPrefix(name, "Panic"),
		strings.HasPrefix(name, "Fatal"),
		strings.HasPrefix(name, "Print"):
		return true
	default:
		return false
	}
}
