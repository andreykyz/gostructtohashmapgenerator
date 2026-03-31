package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// StructField represents a field in a struct.
type StructField struct {
	Name     string // Field name as in source
	Type     string // Go type expression
	TagValue string // Extracted value of the structtomap tag, empty if not present
}

// StructInfo holds information about a parsed struct.
type StructInfo struct {
	Name   string
	Fields []StructField
}

// ParseFile parses a Go source file and returns all structs that have at least one field with the given tag.
func ParseFile(filename, tag string) ([]StructInfo, error) {
	structs, err := ParseAllStructs(filename, tag)
	if err != nil {
		return nil, err
	}
	var tagged []StructInfo
	for _, st := range structs {
		hasTag := false
		for _, f := range st.Fields {
			if f.TagValue != "" {
				hasTag = true
				break
			}
		}
		if hasTag {
			tagged = append(tagged, st)
		}
	}
	return tagged, nil
}

// ParseAllStructs returns all struct definitions in the file, regardless of tags.
// The tag parameter specifies which struct tag to extract values from (e.g., "structtomap").
func ParseAllStructs(filename, tag string) ([]StructInfo, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var structs []StructInfo
	ast.Inspect(f, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}
		// Collect fields
		var fields []StructField
		for _, field := range st.Fields.List {
			if len(field.Names) == 0 {
				// Embedded field, skip for simplicity
				continue
			}
			fieldName := field.Names[0].Name
			typeExpr := string(src[field.Type.Pos()-1 : field.Type.End()-1])
			tagValue := ""
			if field.Tag != nil {
				tagStr := field.Tag.Value
				// Extract tag value for the specified tag
				tagValue = extractTagValue(tagStr, tag)
			}
			fields = append(fields, StructField{
				Name:     fieldName,
				Type:     typeExpr,
				TagValue: tagValue,
			})
		}
		structs = append(structs, StructInfo{
			Name:   ts.Name.Name,
			Fields: fields,
		})
		return true
	})
	return structs, nil
}

// extractTagValue extracts the value for a given key from a struct tag string.
// Example: `structtomap:"name"` -> "name"
func extractTagValue(tag, key string) string {
	tag = strings.Trim(tag, "`")
	for tag != "" {
		// Skip spaces
		tag = strings.TrimSpace(tag)
		// Find colon
		colon := strings.Index(tag, ":")
		if colon < 0 {
			break
		}
		tagKey := tag[:colon]
		tag = tag[colon+1:]
		// Find quoted value
		if len(tag) == 0 || tag[0] != '"' {
			break
		}
		tag = tag[1:]
		quote := strings.Index(tag, `"`)
		if quote < 0 {
			break
		}
		value := tag[:quote]
		tag = tag[quote+1:]
		if tagKey == key {
			return value
		}
	}
	return ""
}
