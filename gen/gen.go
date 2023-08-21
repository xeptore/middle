package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

func init() {
	flag.StringVar(&pkg, "pkg", "github.com/xeptore/middle/v3", "generated file package name")
	flag.StringVar(&filename, "file", "./middle.go", "name of the file to write generated code in")
	flag.IntVar(&n, "n", 26, "number of wares")
}

var alphabets = []string{
	"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func chainStructName(i int) string {
	return fmt.Sprintf("chain%d", i)
}

func chainHandleStructName(i int) string {
	return fmt.Sprintf("chainHandle%d", i)
}

func funcName(i int) string {
	return fmt.Sprintf("Chain%d", i)
}

func genericTypes(i int) []Code {
	return lo.Times(i, func(j int) Code { return Id(alphabets[j]).Any() })
}

func genericTypeParamName(i int) string {
	return strings.ToLower(alphabets[i])
}

func parameterGenericTypes(i int) []Code {
	return lo.Times(i, func(j int) Code { return Id(alphabets[j]) })
}

func fnParams(i int) []Code {
	return lo.Times(i, func(j int) Code {
		return Id(fnName(j+1)).
			Func().
			Params(
				Qual("net/http", "ResponseWriter"),
				Add(Op("*")).Qual("net/http", "Request"),
			).
			Parens(
				List(
					Id(alphabets[j]),
					Error(),
				),
			)
	})
}

func fnName(n int) string {
	return fmt.Sprintf("f%d", n)
}

func handlerFuncType(i int) Code {
	return Id("handler").
		Func().
		Params(
			append(
				[]Code{
					Qual("net/http", "ResponseWriter"),
					Add(Op("*")).Qual("net/http", "Request"),
				},
				parameterGenericTypes(i)...,
			)...,
		).
		Error()
}

var (
	pkg      string
	filename string
	n        int
)

func validateFlags() error {
	if n < 1 || n > 26 {
		return fmt.Errorf("n cannot be < 1 or > 26")
	}
	return nil
}

func main() {
	flag.Parse()
	if err := validateFlags(); nil != err {
		log.Fatalf("provided flags are invalid: %v", err)
	}
	filePathPkg := pkg
	if matches, _ := regexp.MatchString(`/v\d$`, pkg); matches {
		filePathPkg = pkg[:strings.LastIndex(pkg, "/")]
	}
	f := NewFilePath(filePathPkg)
	f.HeaderComment(fmt.Sprintf("Code generated by %s. DO NOT EDIT.", pkg))
	for i := 1; i <= n; i++ {
		f.
			Type().
			Id(chainStructName(i)).
			Types(genericTypes(i)...).
			Struct(fnParams(i)...)

		f.Line()

		f.
			Type().
			Id(chainHandleStructName(i)).
			Types(genericTypes(i)...).
			Struct(
				Id("chain").Id(chainStructName(i)).Types(parameterGenericTypes(i)...),
				handlerFuncType(i),
			)

		f.Comment("Then executes handler once all middleware functions are executed in order as [net/http.Handler]. Chain of functions execution stops if any of the middleware functions returns a non-nil error. It ignores any error returned from handler.")
		f.
			Func().
			Params(Id("chain").Id(chainStructName(i)).Types(parameterGenericTypes(i)...)).
			Id("Then").
			Params(handlerFuncType(i)).
			Id(chainHandleStructName(i)).
			Types(parameterGenericTypes(i)...).
			Block(
				Return(
					Id(chainHandleStructName(i)).
						Types(parameterGenericTypes(i)...).
						Values(
							Id("chain"),
							Id("handler"),
						),
				),
			)

		f.Line()

		f.
			Func().
			Params(Id("chainHandle").Id(chainHandleStructName(i)).Types(parameterGenericTypes(i)...)).
			Id("serveHTTP").
			Params(
				Id("response").Qual("net/http", "ResponseWriter"),
				Id("request").Add(Op("*")).Qual("net/http", "Request"),
			).
			Error().
			Block(
				append(
					lo.Flatten(
						lo.Times(i, func(j int) []Code {
							return []Code{
								List(
									Id(genericTypeParamName(j)),
									Err(),
								).
									Op(":=").
									Id("chainHandle").
									Dot("chain").
									Dot(fnName(j+1)).
									Call(
										Id("response"),
										Id("request"),
									),
								If(
									Nil().
										Op("!=").
										Err(),
								).
									Block(
										Return(Nil()),
									),
							}
						}),
					),
					Return(
						Id("chainHandle").Dot("handler").Call(
							append(
								[]Code{
									Id("response"),
									Id("request"),
								},
								lo.Times(i, func(j int) Code { return Id(genericTypeParamName(j)) })...,
							)...,
						)),
				)...,
			)

		f.Line()

		f.Comment("ServeHTTP satisfies [net/http.Handler].")
		f.
			Func().
			Params(Id("chainHandle").Id(chainHandleStructName(i)).Types(parameterGenericTypes(i)...)).
			Id("ServeHTTP").
			Params(
				Id("response").Qual("net/http", "ResponseWriter"),
				Id("request").Add(Op("*")).Qual("net/http", "Request"),
			).
			Block(Id("_").Op("=").Id("chainHandle").Dot("serveHTTP").Call(Id("response"), Id("request")))

		f.Line()

		f.Comment("Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.")
		f.
			Func().
			Params(Id("chainHandle").Id(chainHandleStructName(i)).Types(parameterGenericTypes(i)...)).
			Id("Finally").
			Params(
				Id("handle").
					Func().
					Params(
						Qual("net/http", "ResponseWriter"),
						Add(Op("*")).Qual("net/http", "Request"),
						Error(),
					),
			).
			Qual("net/http", "HandlerFunc").
			Block(
				Return(
					Func().
						Params(
							Id("response").Qual("net/http", "ResponseWriter"),
							Id("request").Add(Op("*")).Qual("net/http", "Request"),
						).
						Block(
							Err().
								Op(":=").
								Id("chainHandle").Dot("serveHTTP").Call(Id("response"), Id("request")),
							If(Nil().Op("!=").Err()).
								Block(Id("handle").Call(Id("response"), Id("request"), Err())),
						),
				),
			)

		f.Line()

		f.Commentf("%s creates a chain of exactly %d number of function%s that will be executed in order.", funcName(i), i, lo.Ternary(i > 1, "s", ""))
		f.Func().
			Id(funcName(i)).
			Types(genericTypes(i)...).
			Params(fnParams(i)...).
			Id(chainStructName(i)).Types(parameterGenericTypes(i)...).
			Block(
				Return(
					Id(chainStructName(i)).
						Types(parameterGenericTypes(i)...).
						Values(lo.Times(i, func(j int) Code { return Id(fnName(j + 1)) })...),
				),
			)
	}

	var buf bytes.Buffer
	if err := f.Render(&buf); nil != err {
		log.Fatalf("failed to generate code: %v\n", err)
	}
	if err := os.WriteFile(filename, buf.Bytes(), 0644); nil != err {
		log.Fatalf("failed to write generated code to %q: %v\n", filename, err)
	}
}
