package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/loader"
)

var packageName = flag.String("package", "TestInterface", "Destination package name")
var packagePath = flag.String("package-path", "github.com/samirntnx", "Destination package repo path")
var onlyPatch = flag.Bool("only-patch", false, "generate only gopatch")
var packageVar = flag.String("package-var", "Fn", "Global variable that is used by code to access funcs")

func main() {
	var buildTags string
	var includeTests bool
	var verbose bool
	flag.StringVar(&buildTags, "tags", "", "build tags")
	flag.BoolVar(&includeTests, "include-tests", false, "include tests")
	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.Parse()
	prog, _, err := loadProgram(parseBuildTags(buildTags), []string{*packageName}, includeTests)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = printFuncsInProgram(prog, verbose); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func parseBuildTags(tags string) []string {
	var result []string
	split := strings.Split(tags, ",")
	for _, s := range split {
		result = append(result, strings.TrimSpace(s))
	}
	return result
}

func loadProgram(tags, args []string, includeTests bool) (*loader.Program, []string, error) {
	var conf loader.Config
	conf.Build = &build.Default
	conf.Build.BuildTags = append(conf.Build.BuildTags, tags...)
	rest, err := conf.FromArgs(args, includeTests)
	if err != nil {
		return nil, rest, err
	}
	prog, err := conf.Load()
	return prog, rest, err
}

var structs map[string][]*ast.FuncDecl
var allTypes map[string]struct{}
var funcs map[string]*ast.FuncDecl

func printFuncsInProgram(prog *loader.Program, verbose bool) error {
	structs = make(map[string][]*ast.FuncDecl)
	funcs = make(map[string]*ast.FuncDecl)
	allTypes = make(map[string]struct{})

	for _, pkgInfo := range prog.InitialPackages() {
		for _, file := range pkgInfo.Files {
			if err := collectFuncsInFile(file); err != nil {
				return err
			}
		}
	}

	if !*onlyPatch {
		fmt.Printf("package %s\n\n", filepath.Base(*packageName))
		fmt.Printf("type any = interface{}\n\n")
		printInterface()
		printDirectCalls()
	}

	return nil
}

func printMocks() {
	bPackage := filepath.Base(*packageName)
	cPackage := toCamelCase(bPackage)

	fmt.Printf("// Test override the below var to provide mockgen mock of package interface\n")
	fmt.Printf("var Var%sMock %sInterface\n\n", cPackage, cPackage)

	fmt.Printf("// %sMockInit setups the mock for package funcs\n", cPackage)
	fmt.Printf("func %sMockInit() {\n", cPackage)
	for ii, fd := range funcs {

		fields := make([]string, 0)
		for _, f := range fd.Type.Params.List {
			for _, n := range f.Names {
				fields = append(fields, n.Name)
			}
		}

		returnStr := ""
		if fd.Type.Results != nil && len(fd.Type.Results.List) > 0 {
			returnStr = fmt.Sprintf("\treturn Var%sMock.%s(%s)", cPackage, fd.Name.Name, strings.Join(fields, ","))
		} else {
			returnStr = fmt.Sprintf("\tVar%sMock.%s(%s)", cPackage, fd.Name.Name, strings.Join(fields, ","))
		}

		fmt.Printf("\n\tmock%s := mocker.NewMock()\n", ii)
		fmt.Printf("\t"+`mock%s.Patch(%s.%s).AnyTimes().DoAndReturn(
			func(%s) %s{
				if Var%sMock != nil {
					%s
				}
				panic("%s Mock not installed")
			})`+"\n", ii, bPackage, fd.Name.Name, formatFuncParams(fd.Type.Params),
			formatFuncResults(fd.Type.Results), cPackage, returnStr, cPackage)
	}

	// for k := range structs {
	// 	fmt.Printf("\t%sMockInit()\n", toCamelCase(k))
	// }

	fmt.Printf("}\n")

	for k, v := range structs {
		//structName := normalizePackageName(*packageName) + "." + k
		fmt.Printf("// Test override the below var to provide mockgen mock of package interface\n")
		fmt.Printf("var Var%sMock %sInterface\n\n", toCamelCase(k), toCamelCase(k))
		fmt.Printf("// %sMockInit setups the mock for package funcs\n", toCamelCase(k))
		fmt.Printf("func %sMockInit() {\n", toCamelCase(k))

		for ii, fd := range v {
			fields := make([]string, 0)
			for _, f := range fd.Type.Params.List {
				for _, n := range f.Names {
					fields = append(fields, n.Name)
				}
			}

			returnStr := ""
			if fd.Type.Results != nil && len(fd.Type.Results.List) > 0 {
				returnStr = fmt.Sprintf("\treturn Var%sMock.%s(%s)", toCamelCase(k), fd.Name.Name, strings.Join(fields, ","))
			} else {
				returnStr = fmt.Sprintf("\tVar%sMock.%s(%s)", toCamelCase(k), fd.Name.Name, strings.Join(fields, ","))
			}

			fmt.Printf("\n\tmock%d := mocker.NewMock()\n", ii)
			fmt.Printf("\t"+`mock%d.PatchInstance(&%s.%s{}, "%s").AnyTimes().DoAndReturn(
				func(_ *%s.%s, %s) %s{
					if Var%sMock != nil {
						%s
					}
					panic("%s Mock not installed")
				})`+"\n", ii, bPackage, k, fd.Name.Name, bPackage, k, formatFuncParams(fd.Type.Params),
				formatFuncResults(fd.Type.Results), toCamelCase(k), returnStr, k)
		}
		fmt.Printf("}\n")
	}
}

func printDirectCalls() {
	bPackage := filepath.Base(*packageName)
	cPackage := toCamelCase(bPackage)

	fmt.Printf("// Test override the below var to provide mockgen mock of package interface\n")
	fmt.Printf("var Var%sMock %sInterface = nil\n\n", cPackage, cPackage)
	for ii, fd := range funcs {
		fmt.Printf("func %s(%s) %s {\n", ii, formatFuncParams(fd.Type.Params),
			formatFuncResults(fd.Type.Results))

		// fmt.Printf("\tif Var%sMock == nil {\n", cPackage)
		// fmt.Printf("\t\tpanic(\"Var%sMock not initialized\")\n", cPackage)
		// fmt.Printf("\t}\n")

		fields := make([]string, 0)
		for _, f := range fd.Type.Params.List {
			for _, n := range f.Names {
				fields = append(fields, n.Name)
			}
		}

		fmt.Printf("\tif Var%sMock == nil {\n", cPackage)
		if fd.Type.Results != nil && len(fd.Type.Results.List) > 0 {
			fmt.Printf("\t\treturn Var%sMock.%s(%s)\n", cPackage, fd.Name.Name, strings.Join(fields, ","))
		} else {
			fmt.Printf("\t\tVar%sMock.%s(%s)\n", cPackage, fd.Name.Name, strings.Join(fields, ","))
		}
		fmt.Printf("\t}\n")

		if fd.Type.Results != nil && len(fd.Type.Results.List) > 0 {
			fmt.Printf("\t\treturn %s.%s(%s)\n", bPackage, fd.Name.Name, strings.Join(fields, ","))
		} else {
			fmt.Printf("\t\t%s.%s(%s)\n", bPackage, fd.Name.Name, strings.Join(fields, ","))
		}
		fmt.Printf("}\n\n")

	}
}

func normalizePackageName(name string) string {
	return strings.Replace(name, "/", ".", -1)
}

func printGoPatch() {
	bPackage := filepath.Base(*packageName)
	fmt.Printf("@@\n@@\n")
	fmt.Printf("-import \"%s\"\n", *packageName)
	fmt.Printf("+import \"%s\"\n", *packageName)
	fmt.Printf("+import g%s \"%s/%s\"\n", bPackage, *packagePath, *packageName)
	fmt.Printf("\n%s\n", bPackage)

	for _, fd := range funcs {
		fmt.Printf("@@\n@@\n")
		fmt.Printf("-%s.%s(...)\n", bPackage, fd.Name.Name)
		fmt.Printf("+g%s.%s.%s(...)\n\n", bPackage, *packageVar, fd.Name.Name)
	}

	for k := range structs {
		fmt.Printf("@@\n@@\n")
		fmt.Printf("-%s.%s\n", bPackage, k)
		fmt.Printf("+g%s.%s\n\n", bPackage, k)
	}
}

func collectFuncsInFile(file *ast.File) error {
	for _, xdecl := range file.Decls {
		switch decl := xdecl.(type) {
		case *ast.FuncDecl:
			if exported(decl) {
				if decl.Recv != nil {
					if _, ok := structs[structName(decl.Recv.List[0].Type)]; !ok {
						structs[structName(decl.Recv.List[0].Type)] = make([]*ast.FuncDecl, 0)
					}
					structs[structName(decl.Recv.List[0].Type)] = append(structs[structName(decl.Recv.List[0].Type)], decl)
				} else {
					funcs[decl.Name.Name] = decl
				}
			}
		}
	}

	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		node := c.Node()
		switch x := node.(type) {
		case *ast.TypeSpec:
			allTypes[x.Name.Name] = struct{}{}
		}
		return true
	})
	return nil
}

func printInterface() {
	bPackage := filepath.Base(*packageName)
	fmt.Printf("type %sInterface interface {\n", toCamelCase(bPackage))
	for _, decl := range funcs {
		fmt.Printf("\t%s(%s) %s\n", decl.Name.Name,
			formatFuncParams(decl.Type.Params),
			formatFuncResults(decl.Type.Results))
	}
	fmt.Printf("}\n\n")

	for k, v := range structs {
		fmt.Printf("type %sInterface interface {\n", toCamelCase(k))
		for _, decl := range v {
			fmt.Printf("\t%s(%s) %s\n", decl.Name.Name,
				formatFuncParams(decl.Type.Params),
				formatFuncResults(decl.Type.Results))
		}
		fmt.Printf("}\n\n")
	}
}

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

func printWrappers() {
	fmt.Printf("type %sImpl struct {}\n", toCamelCase(*packageName))
	for _, decl := range funcs {
		fmt.Printf("func (recv *%sImpl) %s(%s) %s {\n",
			toCamelCase(*packageName),
			decl.Name.Name,
			formatFuncParams(decl.Type.Params),
			formatFuncResults(decl.Type.Results))
		fields := make([]string, 0)
		for _, f := range decl.Type.Params.List {
			for _, n := range f.Names {
				fields = append(fields, n.Name)
			}
		}
		if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
			fmt.Printf("\treturn %s.%s(%s)\n", *packageName, decl.Name.Name, strings.Join(fields, ","))
		} else {
			fmt.Printf("\t%s.%s(%s)\n", *packageName, decl.Name.Name, strings.Join(fields, ","))
		}

		fmt.Printf("}\n\n")
	}

	for k, v := range structs {
		fmt.Printf("type %sImpl struct {\n", toCamelCase(k))
		fmt.Printf("\t*%s.%s\n", *packageName, k)
		fmt.Printf("}\n\n")
		for _, decl := range v {
			fmt.Printf("func (recv *%sImpl) %s(%s) %s {\n",
				toCamelCase(k),
				decl.Name.Name,
				formatFuncParams(decl.Type.Params),
				formatFuncResults(decl.Type.Results))
			fields := make([]string, 0)
			for _, f := range decl.Type.Params.List {
				for _, n := range f.Names {
					fields = append(fields, n.Name)
				}
			}
			if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
				fmt.Printf("\treturn recv.%s.%s(%s)\n", k, decl.Name.Name, strings.Join(fields, ","))
			} else {
				fmt.Printf("\trecv.%s.%s(%s)\n", k, decl.Name.Name, strings.Join(fields, ","))
			}
			fmt.Printf("}\n\n")
		}
	}
}

func printFuncsInFile(file *ast.File, verbose bool) error {
	for _, xdecl := range file.Decls {
		switch decl := xdecl.(type) {
		case *ast.FuncDecl:
			if exported(decl) {
				if verbose {
					fmt.Println(formatFuncDecl(decl))
				} else {
					fmt.Println(decl.Name.Name)
				}
			}
		}
	}
	return nil
}

func exported(decl *ast.FuncDecl) bool {
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
		return isUpper0(formatType(field.Type)) && isUpper0(decl.Name.Name)
	}
	return isUpper0(decl.Name.Name)
}

func structName(typ ast.Expr) string {
	return typ.(*ast.StarExpr).X.(*ast.Ident).Name
}

func formatType(typ ast.Expr) string {
	bPackage := filepath.Base(*packageName)
	switch t := typ.(type) {
	case nil:
		return ""

	case *ast.Ident:
		// if _, ok := structs[t.Name]; ok {
		// 	return fmt.Sprintf("%sInterface", toCamelCase(t.Name))
		// }
		if _, ok := allTypes[t.Name]; ok {
			return fmt.Sprintf("%s.%s", bPackage, toCamelCase(t.Name))
		}
		return t.Name

	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", formatType(t.X), t.Sel.Name)

	case *ast.StarExpr:
		x := fmt.Sprintf("*%s", formatType(t.X))
		if strings.HasSuffix(x, "Interface") {
			return formatType(t.X)
		}
		// if _, ok := allTypes[formatType(t.X)]; ok {
		// 	return fmt.Sprintf("%s.%s", *packageName, toCamelCase(x))
		// }
		return x

	case *ast.ArrayType:
		// elt := formatType(t.Elt)
		// if _, ok := allTypes[elt]; ok {
		// 	return fmt.Sprintf("[%s]%s.%s", formatType(t.Len), *packageName, toCamelCase(elt))
		// }

		return fmt.Sprintf("[%s]%s", formatType(t.Len), formatType(t.Elt))

	case *ast.Ellipsis:
		return formatType(t.Elt)

	case *ast.FuncType:
		return fmt.Sprintf("func(%s)%s", formatFuncParams(t.Params), formatFuncResults(t.Results))

	case *ast.MapType:
		// key := formatType(t.Key)
		// str := ""
		// if _, ok := allTypes[key]; ok {
		// 	str += fmt.Sprintf("[%s.%s]", *packageName, key)
		// }
		// value := formatType(t.Value)
		// if _, ok := allTypes[value]; ok {
		// 	str += fmt.Sprintf("%s.%s", *packageName, value)
		// }

		// if str != "" {
		// 	return str
		// }
		return fmt.Sprintf("map[%s]%s", formatType(t.Key), formatType(t.Value))
	case *ast.ChanType:
		// FIXME
		panic(fmt.Errorf("unsupported chan type %#v", t))
	case *ast.BasicLit:
		return t.Value
	default:
		panic(fmt.Errorf("unsupported type %#v", t))
	}
}

func formatFields(fields *ast.FieldList) string {
	s := ""
	for i, field := range fields.List {
		for j, name := range field.Names {
			s += name.Name
			if j != len(field.Names)-1 {
				s += ","
			}
			s += " "
		}
		s += formatType(field.Type)
		if i != len(fields.List)-1 {
			s += ", "
		}
	}
	return s
}

func formatFuncParams(fields *ast.FieldList) string {
	return formatFields(fields)
}

func formatFuncResults(fields *ast.FieldList) string {
	s := ""
	if fields != nil {
		s += " "
		if len(fields.List) > 1 {
			s += "("
		}
		s += formatFields(fields)
		if len(fields.List) > 1 {
			s += ")"
		}
	}
	return s
}

func formatFuncDecl(decl *ast.FuncDecl) string {
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
		s += fmt.Sprintf("(%s %s) ", field.Names[0], formatType(field.Type))
	}
	s += fmt.Sprintf("%s(%s)", decl.Name.Name, formatFuncParams(decl.Type.Params))
	s += formatFuncResults(decl.Type.Results)
	return s
}
