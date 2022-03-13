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
		(*ast.CallExpr)(nil),
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
			sl := make(map[int]int)
			for _, l := range n.Lhs {
				switch ln := l.(type) {
				case *ast.Ident:
					typ := pass.TypesInfo.TypeOf(l)
					if ln.Name == "_" {
						if isErrorType(typ) {
							pass.Reportf(n.Pos(), "blank error")
							return false
						}

						if types.Identical(typ, nil) {
							for _, r := range n.Rhs {
								if rl, ok := r.(*ast.CallExpr); ok {
									if ro, ok := rl.Fun.(*ast.Ident); ok {
										if rd, ok := ro.Obj.Decl.(*ast.FuncDecl); ok {
											for i, t := range rd.Type.Results.List {
												typ2 := pass.TypesInfo.TypeOf(t.Type)
												if isErrorType(typ2) {
													sl[i]++
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			for i := range sl {
				pass.Reportf(n.Lhs[i].Pos(), "blank error")
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
