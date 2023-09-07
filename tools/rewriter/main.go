package main

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	"go/ast"

	"go/build"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

var structs map[string](map[string][](*ast.FuncDecl))
var funcs map[string](map[string]*ast.FuncDecl)
var allTypes map[string](map[string]struct{})

func toCamelCase(input string) string {
	// Split the input string into words
	input = strings.Replace(input, "*", "", 1)
	input = strings.Replace(input, "/", "", -1)
	words := strings.Fields(input)

	// Iterate over the words and capitalize the first letter of each word
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	// Join the words and convert the first letter to lowercase
	camelCase := strings.Join(words, "")
	camelCase = string(unicode.ToUpper(rune(camelCase[0]))) + camelCase[1:]

	return camelCase
}

func loadPackage(pkg string) error {
	log.Printf("Loading package: %q", pkg)
	var conf loader.Config
	conf.Build = &build.Default
	_, err := conf.FromArgs([]string{pkg}, false)
	if err != nil {
		return err
	}

	prg, err := conf.Load()
	if err != nil {
		return err
	}

	allTypes[pkg] = make(map[string]struct{})
	structs[pkg] = make(map[string][]*ast.FuncDecl)
	funcs[pkg] = make(map[string]*ast.FuncDecl)

	for _, pkgInfo := range prg.InitialPackages() {
		for _, file := range pkgInfo.Files {
			if err := collectFuncsInFile(pkg, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func exported(pkg string, decl *ast.FuncDecl) bool {
	isUpper0 := func(s string) bool {
		if strings.HasPrefix(s, "*") {
			return unicode.IsUpper([]rune(s)[1])
		}
		return unicode.IsUpper([]rune(s)[0])
	}
	if decl.Recv != nil {
		if len(decl.Recv.List) != 1 {
			panic(fmt.Errorf("strange receiver for %s: %#v", decl.Name.Name, decl.Recv))
		}
		field := decl.Recv.List[0]
		return isUpper0(formatType(pkg, field.Type)) && isUpper0(decl.Name.Name)
	}
	return isUpper0(decl.Name.Name)
}

func structName(typ ast.Expr) string {
	return typ.(*ast.StarExpr).X.(*ast.Ident).Name
}

func formatType(pkg string, typ ast.Expr) string {
	bPackage := filepath.Base(pkg)
	switch t := typ.(type) {
	case nil:
		return ""

	case *ast.Ident:
		if _, ok := structs[t.Name]; ok {
			return fmt.Sprintf("%sInterface", toCamelCase(t.Name))
		}
		if _, ok := allTypes[t.Name]; ok {
			return fmt.Sprintf("%s.%s", bPackage, toCamelCase(t.Name))
		}
		return t.Name

	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", formatType(pkg, t.X), t.Sel.Name)

	case *ast.StarExpr:
		x := fmt.Sprintf("*%s", formatType(pkg, t.X))
		if strings.HasSuffix(x, "Interface") {
			return formatType(pkg, t.X)
		}
		// if _, ok := allTypes[formatType(pkg,t.X)]; ok {
		// 	return fmt.Sprintf("%s.%s", pkg, toCamelCase(x))
		// }
		return x

	case *ast.ArrayType:
		// elt := formatType(pkg,t.Elt)
		// if _, ok := allTypes[elt]; ok {
		// 	return fmt.Sprintf("[%s]%s.%s", formatType(pkg,t.Len), pkg, toCamelCase(elt))
		// }

		return fmt.Sprintf("[%s]%s", formatType(pkg, t.Len), formatType(pkg, t.Elt))

	case *ast.Ellipsis:
		return formatType(pkg, t.Elt)

	case *ast.FuncType:
		return fmt.Sprintf("func(%s)%s", formatFuncParams(pkg, t.Params), formatFuncResults(pkg, t.Results))

	case *ast.MapType:
		// key := formatType(pkg,t.Key)
		// str := ""
		// if _, ok := allTypes[key]; ok {
		// 	str += fmt.Sprintf("[%s.%s]", pkg, key)
		// }
		// value := formatType(pkg,t.Value)
		// if _, ok := allTypes[value]; ok {
		// 	str += fmt.Sprintf("%s.%s", pkg, value)
		// }

		// if str != "" {
		// 	return str
		// }
		return fmt.Sprintf("map[%s]%s", formatType(pkg, t.Key), formatType(pkg, t.Value))
	case *ast.ChanType:
		// FIXME
		panic(fmt.Errorf("unsupported chan type %#v", t))
	case *ast.BasicLit:
		return t.Value
	default:
		panic(fmt.Errorf("unsupported type %#v", t))
	}
}

func formatFields(pkg string, fields *ast.FieldList) string {
	s := ""
	for i, field := range fields.List {
		for j, name := range field.Names {
			s += name.Name
			if j != len(field.Names)-1 {
				s += ","
			}
			s += " "
		}
		s += formatType(pkg, field.Type)
		if i != len(fields.List)-1 {
			s += ", "
		}
	}
	return s
}

func formatFuncParams(pkg string, fields *ast.FieldList) string {
	return formatFields(pkg, fields)
}

func formatFuncResults(pkg string, fields *ast.FieldList) string {
	s := ""
	if fields != nil {
		s += " "
		if len(fields.List) > 1 {
			s += "("
		}
		s += formatFields(pkg, fields)
		if len(fields.List) > 1 {
			s += ")"
		}
	}
	return s
}

func formatFuncDecl(pkg string, decl *ast.FuncDecl) string {
	s := "func "
	if decl.Recv != nil {
		if len(decl.Recv.List) != 1 {
			panic(fmt.Errorf("strange receiver for %s: %#v", decl.Name.Name, decl.Recv))
		}
		field := decl.Recv.List[0]
		if len(field.Names) == 0 {
			// function definition in interface (ignore)
			return ""
		} else if len(field.Names) != 1 {
			panic(fmt.Errorf("strange receiver field for %s: %#v", decl.Name.Name, field))
		}
		s += fmt.Sprintf("(%s %s) ", field.Names[0], formatType(pkg, field.Type))
	}
	s += fmt.Sprintf("%s(%s)", decl.Name.Name, formatFuncParams(pkg, decl.Type.Params))
	s += formatFuncResults(pkg, decl.Type.Results)
	return s
}

func collectFuncsInFile(pkg string, file *ast.File) error {
	for _, xdecl := range file.Decls {
		switch decl := xdecl.(type) {
		case *ast.FuncDecl:
			if exported(pkg, decl) {
				if decl.Recv != nil {
					if _, ok := structs[structName(decl.Recv.List[0].Type)]; !ok {
						structs[pkg][structName(decl.Recv.List[0].Type)] = make([]*ast.FuncDecl, 0)
					}
					structs[pkg][structName(decl.Recv.List[0].Type)] = append(structs[pkg][structName(decl.Recv.List[0].Type)], decl)
				} else {
					funcs[pkg][decl.Name.Name] = decl
				}
			}
		}
	}

	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		node := c.Node()
		switch x := node.(type) {
		case *ast.TypeSpec:
			allTypes[pkg][x.Name.Name] = struct{}{}
		}
		return true
	})
	return nil
}

func main() {
	// Define the Go source code file you want to modify.
	filePath := os.Args[1]
	allTypes = make(map[string]map[string]struct{})
	structs = make(map[string]map[string][]*ast.FuncDecl)
	funcs = make(map[string]map[string]*ast.FuncDecl)

	spec := os.Args[2]
	changes := make(map[string]string)
	for _, v := range strings.Split(spec, ",") {
		ts := strings.Split(v, "=")
		changes[ts[0]] = ts[1]
		loadPackage(ts[0])
	}

	// Create a new token set and parser.
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		changed := false

		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {

			fs := token.NewFileSet()
			file, err := parser.ParseFile(fs, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatalf("Error parsing file: %v", err)
			}

			// Replace the old package name with the new package name in the import statements.
			rewriter := func(cursor *astutil.Cursor) bool {
				node := cursor.Node()
				if node == nil {
					return true
				}

				// Check if it's an import declaration.
				if importSpec, ok := node.(*ast.ImportSpec); ok {
					path := strings.Trim(importSpec.Path.Value, `"`)
					if val, ok := changes[path]; ok {
						importSpec.Path.Value = fmt.Sprintf(`"%s"`, val)
						importSpec.Name = &ast.Ident{Name: fmt.Sprintf("g%s", filepath.Base(path))}
						changed = true
					}
				}

				if selectorExpr, ok := node.(*ast.SelectorExpr); ok {
					if ident, ok := selectorExpr.X.(*ast.Ident); ok {
						if _, ok := changes[ident.Name]; ok {
							if findSymbolInPkgs(ident.Name, selectorExpr.Sel.Name) {
								ident.Name = fmt.Sprintf("g%s", filepath.Base(ident.Name))
								changed = true
							}
						}
					}
				}

				return true
			}

			astutil.Apply(file, rewriter, nil)

			if !changed {
				return nil
			}

			fmt.Printf("Patching %q file\n", path)
			// Write the modified AST back to the file.
			outputFile, err := os.Create(path)
			if err != nil {
				log.Fatalf("Error creating output file: %v", err)
			}
			defer outputFile.Close()

			if err := printer.Fprint(outputFile, fs, file); err != nil {
				log.Fatalf("Error writing to output file: %v", err)
			}

			// Run goimports on the output file.
			cmd := exec.Command("goimports", "-w", path)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("Error running goimports: %v", err)
			}
		}
		return nil
	})

	// Parse the Go source file.
}

func findSymbolInPkgs(pkg, s string) bool {
	if _, ok := structs[pkg][s]; ok {
		return true
	}
	if _, ok := funcs[pkg][s]; ok {
		return true
	}
	if _, ok := allTypes[pkg][s]; ok {
		return true
	}
	return false
}
