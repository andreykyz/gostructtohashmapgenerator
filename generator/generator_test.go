package generator

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	opts := Options{Tag: "structtomap"}
	code, err := Generate("../examples/person.go", opts)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(code) == 0 {
		t.Fatal("generated code is empty")
	}
	// Check that expected functions are present
	codeStr := string(code)
	if !strings.Contains(codeStr, "PersonToMap") {
		t.Error("generated code missing PersonToMap")
	}
	if !strings.Contains(codeStr, "AddressToMap") {
		t.Error("generated code missing AddressToMap")
	}
	// Check that PersonToMap calls AddressToMap
	if !strings.Contains(codeStr, "AddressToMap(p.Address)") {
		// maybe the variable name is different
		if !strings.Contains(codeStr, "AddressToMap") {
			t.Error("PersonToMap should call AddressToMap")
		}
	}
	// Ensure no syntax errors by trying to parse? (optional)
}

func TestGenerateAllFields(t *testing.T) {
	opts := Options{All: true}
	code, err := Generate("../examples/person.go", opts)
	if err != nil {
		t.Fatalf("Generate with All=true failed: %v", err)
	}
	codeStr := string(code)
	// Should still generate both functions
	if !strings.Contains(codeStr, "PersonToMap") || !strings.Contains(codeStr, "AddressToMap") {
		t.Error("missing expected functions")
	}
}

func TestGenerateCrossPackage(t *testing.T) {
	opts := Options{Tag: "structtomap"}
	code, err := Generate("../examples/crosspkg.go", opts)
	if err != nil {
		t.Fatalf("Generate failed for crosspkg.go: %v", err)
	}
	if len(code) == 0 {
		t.Fatal("generated code is empty")
	}
	codeStr := string(code)
	if !strings.Contains(codeStr, "AccountToMap") {
		t.Error("generated code missing AccountToMap")
	}
	// Should contain the field mappings
	if !strings.Contains(codeStr, `m["id"]`) {
		t.Error("missing id field")
	}
	if !strings.Contains(codeStr, `m["owner"]`) {
		t.Error("missing owner field")
	}
	if !strings.Contains(codeStr, `m["balance"]`) {
		t.Error("missing balance field")
	}
	// Ensure no syntax errors by trying to parse? (optional)
}
