package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

var alphabets = []string{
	"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func init() {
	flag.StringVar(&pkg, "pkg", "github.com/xeptore/middle", "generated file package name")
	flag.StringVar(&filename, "file", "./middle.go", "name of the file to write generated code in")
	flag.IntVar(&n, "n", 26, "number of wares")
}

func structName(i int) string {
	return fmt.Sprintf("mw%d", i)
}

func funcName(i int) string {
	return fmt.Sprintf("Chain%d", i)
}

func genericTypes(i int) []Code {
	return lo.Times(i, func(j int) Code { return Id(alphabets[j]).Any() })
}

func parameterGenericTypes(i int) []Code {
	return lo.Times(i, func(j int) Code { return Id(alphabets[j]) })
}

func fnParams(i int) []Code {
	return lo.Times(i, func(j int) Code {
		return Id(fnName(j+1)).Func().Params(Qual("net/http", "ResponseWriter"), Add(Op("*")).Qual("net/http", "Request")).Parens(List(Id(alphabets[j]), Error()))
	})
}

func fnName(n int) string {
	return fmt.Sprintf("f%d", n)
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
	f := NewFilePath(pkg)
	f.HeaderComment("Code generated by middle. DO NOT EDIT.")
	for i := 1; i <= n; i++ {
		f.
			Type().
			Id(structName(i)).
			Types(genericTypes(i)...).
			Struct(fnParams(i)...)
		f.Commentf("Then executes handler once all middleware functions are executed in order. Chain of functions execution stops if any of the middleware functions returns a non-nil error.")
		f.
			Func().
			Params(Id("middleware").Id(structName(i)).Types(parameterGenericTypes(i)...)).
			Id("Then").
			Params(
				Id("handler").
					Func().
					Params(
						append(
							[]Code{
								Qual("net/http", "ResponseWriter"),
								Add(Op("*")).Qual("net/http", "Request"),
							},
							parameterGenericTypes(i)...,
						)...,
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
							append(
								lo.Flatten(
									lo.Times(i, func(j int) []Code {
										return []Code{
											List(
												Id(strings.ToLower(alphabets[j])),
												Err(),
											).
												Op(":=").
												Id("middleware").
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
													Return(),
												),
										}
									}),
								),
								Id("handler").Call(
									append(
										[]Code{
											Id("response"),
											Id("request"),
										},
										lo.Times(i, func(j int) Code { return Id(strings.ToLower(alphabets[j])) })...,
									)...,
								),
							)...,
						),
				),
			)
		f.Line()
		f.Commentf("%s creates a chain of exactly %d number of function%s that will be executed in order.", funcName(i), i, lo.Ternary(i > 1, "s", ""))
		f.Func().
			Id(funcName(i)).
			Types(genericTypes(i)...).
			Params(fnParams(i)...).
			Id(structName(i)).Types(parameterGenericTypes(i)...).
			Block(
				Return(
					Id(structName(i)).
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