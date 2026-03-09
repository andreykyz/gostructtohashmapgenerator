# Go Struct to HashMap Generator

A Go tool that automatically generates code to convert Go structs into `map[string]any` (hashmap) based on struct tags. Supports nested structs and customizable field mapping.

## Features

- **Struct Tag Detection**: Uses a custom tag `structtomap` to indicate which fields should be included in the map.
- **Nested Struct Support**: Recursively converts nested struct fields into nested maps, including cross‑package types.
- **Code Generation**: Produces ready‑to‑use Go functions that perform the conversion at runtime. **One function per struct** – even if the same struct appears nested multiple times, a single conversion function is generated and reused.
- **Import Preservation**: Automatically copies import statements from the source file into the generated code.
- **CLI & Library**: Can be used as a command‑line tool or imported as a library.
- **Zero Dependencies**: Pure Go, no external dependencies.

## Installation

```bash
go install github.com/andrey/gostructtohashmapgenerator/cmd/gostructtohashmap@latest
```

Or clone the repository and build locally:

```bash
git clone https://github.com/andrey/gostructtohashmapgenerator.git
cd gostructtohashmapgenerator
go build ./cmd/gostructtohashmap
```

## Usage

### 1. Tag Your Struct

Add the `structtomap` tag to fields you want to include. The tag value can be used to specify a custom key name; if empty, the field name is used.

```go
type Person struct {
    Name    string `structtomap:"name"`
    Age     int    `structtomap:"age"`
    Address struct {
        City    string `structtomap:"city"`
        Country string `structtomap:"country"`
    } `structtomap:"address"`
}
```

### 2. Generate Conversion Code

Run the generator on your source file:

```bash
gostructtohashmap -input=person.go -output=person_map.go
```

This will create a file `person_map.go` containing a function `PersonToMap(p Person) map[string]any`. If `Address` is also a struct, a separate `AddressToMap` function will be generated and called internally.

### 3. Use the Generated Function

```go
p := Person{
    Name: "Alice",
    Age:  30,
    Address: Address{City: "Moscow", Country: "Russia"},
}
m := PersonToMap(p)
fmt.Printf("%#v", m)
// Output: map[string]any{"name":"Alice", "age":30, "address":map[string]any{"city":"Moscow", "country":"Russia"}}
```

## Advanced Options

- **Exclude Untagged Fields**: By default, only tagged fields are included. Use `-all` to include all fields.
- **Custom Tag Name**: Use `-tag=myTag` to look for a different tag.

## Library Usage

You can also use the generator programmatically:

```go
import "github.com/andrey/gostructtohashmapgenerator/generator"

code, err := generator.Generate("person.go", generator.Options{
    Tag: "structtomap",
    All: false,
})
```

## Project Structure

```
gostructtohashmapgenerator/
├── cmd/gostructtohashmap/   # CLI entry point
├── generator/               # Core generation logic (package)
├── parser/                  # Go source parsing (package)
├── examples/                # Example structs and usage
└── (other auxiliary files)
```

## Examples

See the [`examples/`](examples/) directory for ready‑to‑run examples:

- `person.go` – simple nested struct
- `event.go` – struct with slices and pointers
- `crosspkg.go` – using a type from another package (`examples/models/user.go`)

Each example includes a generated `*_map.go` file that you can inspect.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

MIT