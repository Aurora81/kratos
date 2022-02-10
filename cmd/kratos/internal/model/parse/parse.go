package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func findStructByName(f ast.Node, name string) *ast.TypeSpec {
	var r *ast.TypeSpec = nil

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == name {
				if _, ok := x.Type.(*ast.StructType); ok {
					r = x
					return false
				}
			}
		}
		return true
	})

	return r
}

func Parse(filename string, modelName string) (fields []string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	x := findStructByName(f, modelName)
	if x != nil {
		s := x.Type.(*ast.StructType)
		for _, field := range s.Fields.List {
			for _, name := range field.Names {
				fields = append(fields, name.String())
			}
		}
	}

	return
}