package noniljson

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"
	"strings"
)

var Analyzer = &analysis.Analyzer{
	Name:     "noniljson",
	Doc:      "checks that nullable fields in structs used for JSON marshaling use 'omitempty' or are otherwise correctly handled",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func init() {
	Analyzer.Flags.BoolVar(&ignoreGeneratedFiles, "ignore-generated", false, "ignore files with generated comment in the header")
}

func run(pass *analysis.Pass) (interface{}, error) {
	if ignoreGeneratedFiles {
		for _, f := range pass.Files {
			if isGeneratedFile(f.Comments) {
				return nil, nil
			}
		}
	}
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		typeSpec := n.(*ast.TypeSpec)
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return // Not a struct type
		}

		for _, field := range structType.Fields.List {
			if field.Tag == nil {
				continue // No JSON tag present
			}
			tagValue := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
			jsonTag, ok := tagValue.Lookup("json")
			if !ok || strings.Contains(jsonTag, "-") {
				continue // Not marshaled to JSON or explicitly ignored
			}

			fieldType := pass.TypesInfo.TypeOf(field.Type)
			if isNullableType(fieldType) && !strings.Contains(jsonTag, "omitempty") {
				for _, fieldName := range field.Names {
					pass.Reportf(field.Pos(), "nullable field '%s' in struct '%s' must include 'omitempty' in its json tag to avoid marshaling as null", fieldName.Name, typeSpec.Name.Name)
				}
			}
		}
	})

	return nil, nil
}
