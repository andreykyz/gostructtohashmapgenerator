package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/andrey/gostructtohashmapgenerator/generator"
)

func main() {
	var (
		input   = flag.String("input", "", "Input Go source file containing structs")
		output  = flag.String("output", "", "Output file for generated code (default: <input>_map.go)")
		tag     = flag.String("tag", "structtomap", "Struct tag to look for")
		all     = flag.Bool("all", false, "Include fields without the tag")
		reverse = flag.Bool("reverse", false, "Generate reverse conversion functions (MapTo...)")
	)
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Error: -input is required")
		flag.Usage()
		os.Exit(1)
	}
	if *output == "" {
		base := filepath.Base(*input)
		ext := filepath.Ext(base)
		name := base[:len(base)-len(ext)]
		*output = filepath.Join(filepath.Dir(*input), name+"_map.go")
	}

	opts := generator.Options{
		Tag:     *tag,
		All:     *all,
		Reverse: *reverse,
	}
	code, err := generator.Generate(*input, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Generation failed: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(*output, code, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write output: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s successfully.\n", *output)
}
