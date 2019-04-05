package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var templateData = `package {{.Package}}

/*
NOTE! This file is auto generated from another file - DO NOT EDIT!

This file contents has it's source in {{.Filename}}. All structs with only one
field and the suffix 'Type' is being added here. The difference is that the
field XML tag won't have a namespace.
*/

{{ range $t := .Types -}}
// {{$t.StructName}} represents a namespace agnostic version of {{$t.OriginalStructName}}
type {{$t.StructName}} struct {
	{{$t.FieldName}} {{$t.FieldType}} {{$t.FieldTag}}
}

{{ end -}}
`

// typeData holds all the stucts with type suffix that should be re-generated
// with a trimmed XML tag.
type typeData struct {
	Package  string
	Filename string
	Types    []typeStruct
}

type typeStruct struct {
	OriginalStructName string
	StructName         string
	FieldName          string
	FieldType          string
	FieldTag           string
}

func main() {
	var (
		args  []string
		help  bool
		files = []string{}
	)

	flag.BoolVar(&help, "h", false, "Show this help text")
	flag.BoolVar(&help, "help", false, "")

	flag.Parse()

	if help {
		showHelp()

		return
	}

	args = flag.Args()
	if len(args) == 0 {
		args = []string{"./types/..."}
	}

	for _, f := range args {
		if strings.HasSuffix(f, "/...") {
			dir, _ := filepath.Split(f)

			files = append(files, expandGoWildcard(dir)...)

			continue
		}

		// Skip files already auto-generated.
		if strings.HasSuffix(f, "auto_generated.go") {
			continue
		}

		if _, err := os.Stat(f); err == nil {
			files = append(files, f)
		}
	}

	for _, f := range files {
		Process(f)
	}

	return
}

func Process(filename string) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, fileData, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	typeStructs := typeData{
		Package:  file.Name.Name,
		Filename: filename,
		Types:    []typeStruct{},
	}

	// Find all structs defined with only one field and has the suffix "Type"
	for _, x := range file.Decls {
		v, ok := x.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, s := range v.Specs {
			spec, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structName := spec.Name.Name

			if !strings.HasSuffix(structName, "Type") {
				continue
			}

			structType, ok := spec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			if len(structType.Fields.List) != 1 {
				continue
			}

			firstFieldInStruct := structType.Fields.List[0]

			typeStructs.Types = append(typeStructs.Types, typeStruct{
				OriginalStructName: spec.Name.Name,
				FieldName:          firstFieldInStruct.Names[0].Name,
				FieldType:          firstFieldInStruct.Type.(*ast.Ident).Name,
				FieldTag:           firstFieldInStruct.Tag.Value,
			})
		}
	}

	createFile(typeStructs)
}

func createFile(data typeData) {
	for i, d := range data.Types {
		tagParts := strings.Split(d.FieldTag, "\"")
		middleParts := strings.Split(tagParts[1], " ")

		data.Types[i].FieldTag = fmt.Sprintf("`xml:\"%s\"`", middleParts[1])
		data.Types[i].StructName = fmt.Sprintf("%sIn", d.OriginalStructName)
	}

	tmpl := template.Must(template.New("").Parse(templateData))
	buf := bytes.Buffer{}

	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	fileBytes, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	dir, file := filepath.Split(data.Filename)
	filenameParts := strings.Split(file, ".")
	newFilename := fmt.Sprintf("%s_auto_generated.go", strings.Join(filenameParts[:len(filenameParts)-1], "."))
	newFilepath := filepath.Join(dir, newFilename)

	_ = ioutil.WriteFile(newFilepath, fileBytes, 0644)
	fmt.Printf("Generated file: %s\n", newFilename)
}

func expandGoWildcard(root string) []string {
	foundFiles := []string{}

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Skip files already auto-generated.
		if strings.HasSuffix(info.Name(), "auto_generated.go") {
			return nil
		}

		// Only append go files
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		foundFiles = append(foundFiles, path)

		return nil
	})

	return foundFiles
}

func showHelp() {
	helpText := `Usage: type-generator <file> [files...]

Will default to all files in ./types

Flags:`

	fmt.Println(helpText)
	flag.PrintDefaults()
}
