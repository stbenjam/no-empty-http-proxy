package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "noemptyhttpproxy",
	Doc:      "Checks for lack of Proxy setting on HTTP Transport",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		compLit := node.(*ast.CompositeLit)

		if sel, ok := compLit.Type.(*ast.SelectorExpr); ok {
			gopkg, ok := sel.X.(*ast.Ident)
			if !ok {
				return
			}

			if sel.Sel.Name == "Transport" && gopkg.Name == "http" {
				foundProxy := false
				for _, element := range compLit.Elts {
					if kvExpr, ok := element.(*ast.KeyValueExpr); ok {
						if ident, ok := kvExpr.Key.(*ast.Ident); ok && ident.Name == "Proxy" {
							foundProxy = true
						}
					}
				}

				if !foundProxy {
					pass.Reportf(node.Pos(), "http.Transport should set Proxy; http.ProxyFromEnvironment is typical, but you may explicitly set to nil")
				}
			}
		}

	})

	return nil, nil
}
