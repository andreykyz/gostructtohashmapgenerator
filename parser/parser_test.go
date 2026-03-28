package parser

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	structs, err := ParseFile("../examples/person.go", "structtomap")
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	if len(structs) != 2 {
		t.Errorf("expected 2 structs, got %d", len(structs))
	}
	// Check Person
	var person *StructInfo
	var address *StructInfo
	for i := range structs {
		if structs[i].Name == "Person" {
			person = &structs[i]
		}
		if structs[i].Name == "Address" {
			address = &structs[i]
		}
	}
	if person == nil {
		t.Fatal("Person struct not found")
	}
	if address == nil {
		t.Fatal("Address struct not found")
	}
	// Person fields
	if len(person.Fields) != 3 {
		t.Errorf("Person expected 3 fields, got %d", len(person.Fields))
	}
	// Check field tags
	expectedTags := map[string]string{
		"Name":    "name",
		"Age":     "age",
		"Address": "address",
	}
	for _, f := range person.Fields {
		expected, ok := expectedTags[f.Name]
		if !ok {
			t.Errorf("unexpected field %s", f.Name)
			continue
		}
		if f.TagValue != expected {
			t.Errorf("field %s tag value mismatch: got %q, want %q", f.Name, f.TagValue, expected)
		}
	}
	// Address fields
	if len(address.Fields) != 2 {
		t.Errorf("Address expected 2 fields, got %d", len(address.Fields))
	}
}

func TestParseAllStructs(t *testing.T) {
	structs, err := ParseAllStructs("../examples/person.go", "structtomap")
	if err != nil {
		t.Fatalf("ParseAllStructs failed: %v", err)
	}
	if len(structs) != 2 {
		t.Errorf("expected 2 structs, got %d", len(structs))
	}
}
