package blankerr

import (
	"go/ast"
	"go/types"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "blankerr is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "blankerr",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var generatedPattern = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`)

func isGenerated(f *ast.File) bool {
	for _, c := range f.Comments {
		for _, l := range c.List {
			if generatedPattern.MatchString(l.Text) {
				return true
			}
		}
	}

	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.File)(nil),
		(*ast.Ident)(nil),
	}

	inspect.Nodes(nodeFilter, func(n ast.Node, push bool) bool {
		if !push {
			return false
		}

		switch n := n.(type) {
		case *ast.File:
			return !isGenerated(n)

		case *ast.AssignStmt:
			for _, l := range n.Lhs {
				switch n := l.(type) {
				case *ast.Ident:
					typ := pass.TypesInfo.TypeOf(l)
					if n.Name == "_" {
						if isErrorType(typ) {
							pass.Reportf(n.Pos(), "blank error")
						}
					}
				}
			}
			return false
		}

		return false
	})

	return nil, nil
}

func isErrorType(typ types.Type) bool {
	errType := types.Universe.Lookup("error").Type()
	if types.Identical(typ, errType) {
		return true
	}

	if types.Implements(typ, errType.Underlying().(*types.Interface)) {
		return true
	}

	return false
}
