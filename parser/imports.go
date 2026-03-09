package parser

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

// ImportInfo holds information about an import.
type ImportInfo struct {
	Alias string // alias (e.g., "models" or empty if none)
	Path  string // import path (e.g., "github.com/andrey/gostructtohashmapgenerator/examples/models")
}

// ParseImports returns imports from a Go source file.
func ParseImports(filename string) ([]ImportInfo, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var imports []ImportInfo
	for _, imp := range f.Imports {
		info := ImportInfo{
			Path: strings.Trim(imp.Path.Value, `"`),
		}
		if imp.Name != nil {
			info.Alias = imp.Name.Name
		}
		imports = append(imports, info)
	}
	return imports, nil
}
