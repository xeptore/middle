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
	flag.StringVar(&pkg, "pkg", "github.com/xeptore/middle/v5", "generated file package name")
	flag.StringVar(&filename, "file", "./middle.go", "name of the file to write generated code in")
	flag.IntVar(&n, "n", 27, "maximum generated number of chains")
	flag.BoolVar(&noHeader, "no-header", false, "do not generate GENERATED header comment")
}

var alphabets = []string{
	"A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func chainStructName(i int) string {
	return fmt.Sprintf("ChainHandler%d", i)
}

func factoryFuncName(i int) string {
	return fmt.Sprintf("Chain%d", i)
}

func genericTypes(i int) []Code {
	return lo.Times(i-1, func(j int) Code { return Id(alphabets[j]).Any() })
}

func genericTypeParamName(i int) string {
	return strings.ToLower(alphabets[i])
}

func parameterGenericTypes(i int) []Code {
	return lo.Times(i-1, func(j int) Code { return Id(alphabets[j]) })
}

func fnParams(i int) []Code {
	return append(
		lo.Times(i-1, func(j int) Code {
			return Id(fnName(j + 1)).
				Func().
				Params(
					append(
						[]Code{
							Qual("net/http", "ResponseWriter"),
							Add(Op("*")).Qual("net/http", "Request"),
							Add(Op("*")).Qual(sentryPkgQualPath, "Hub"),
							Add(Op("*")).Qual(sentryPkgQualPath, "Span"),
						},
						lo.Times(j, func(j int) Code { return Id(alphabets[j]) })...,
					)...,
				).
				Parens(
					List(
						Id(alphabets[j]),
						Error(),
					),
				)
		}),
		Id(fnName(i)).
			Func().
			Params(
				append(
					[]Code{
						Qual("net/http", "ResponseWriter"),
						Add(Op("*")).Qual("net/http", "Request"),
						Add(Op("*")).Qual(sentryPkgQualPath, "Hub"),
						Add(Op("*")).Qual(sentryPkgQualPath, "Span"),
					},
					lo.Times(i-1, func(j int) Code { return Id(alphabets[j]) })...,
				)...,
			).
			Error(),
	)
}

func fnName(n int) string {
	return fmt.Sprintf("f%d", n)
}

const sentryPkgQualPath = "github.com/getsentry/sentry-go"

var (
	pkg      string
	filename string
	n        int
	noHeader bool
)

func validateFlags() error {
	if n < 1 || n > 27 {
		return fmt.Errorf("n cannot be < 1 or > 27")
	}
	return nil
}

func sentryOnRequestPanicDeferCallbackBlock() []Code {
	return []Code{
		Err().Op(":=").Recover(),
		If(Nil().Op("==").Err()).Block(Return()),
		Id("eventID").Op(":=").Id("hub").Dot("RecoverWithContext").Call(
			Qual("context", "WithValue").Call(Id("request").Dot("Context").Call(), Qual(sentryPkgQualPath, "RequestContextKey"), Id("request")),
			Err(),
		),
		If(Nil().Op("!=").Id("eventID")).Block(
			Id("hub").Dot("Flush").Call(Lit(2).Op("*").Qual("time", "Second")),
		),
		Id("transaction").Dot("Status").Op("=").Qual(sentryPkgQualPath, "HTTPtoSpanStatus").Call(Qual("net/http", "StatusInternalServerError")),
		Id("transaction").Dot("SetTag").Call(Lit("kind"), Lit("panic")),
		Id("hub").Dot("Scope").Call().Dot("SetLevel").Call(Qual(sentryPkgQualPath, "LevelFatal")),
	}
}

func sentryOnRequestPanicDeferCallbackBlockFinally() []Code {
	return append(
		sentryOnRequestPanicDeferCallbackBlock(),
		Switch(Id("v").Op(":=").Err().Assert(Type())).Block(
			Case(Error()).Block(
				Id("catch").Call(Id("response"), Id("request"), Id("hub"), Id("transaction"), Id("v")),
			),
			Case(String()).Block(
				Id("catch").Call(Id("response"), Id("request"), Id("hub"), Id("transaction"), Qual("errors", "New").Call(Id("v"))),
			),
			Case(Qual("fmt", "Stringer")).Block(
				Id("catch").Call(Id("response"), Id("request"), Id("hub"), Id("transaction"), Qual("errors", "New").Call(Qual("fmt", "Sprintf").Call(Lit("%s"), Id("v")))),
			),
		),
	)
}

func sentryOnRequest(panicDeferCallbackBlock []Code) []Code {
	return []Code{
		Id("ctx").Op(":=").Id("request").Dot("Context").Call(),
		Id("hub").Op(":=").Qual(sentryPkgQualPath, "CurrentHub").Call().Dot("Clone").Call(),
		Id("ctx").Op("=").Qual(sentryPkgQualPath, "SetHubOnContext").Call(Id("ctx"), Id("hub")),
		Id("options").Op(":=").Index().Qual(sentryPkgQualPath, "SpanOption").Values(
			Qual(sentryPkgQualPath, "WithOpName").Call(Lit("http.server")),
			Qual(sentryPkgQualPath, "ContinueFromRequest").Call(Id("request")),
			Qual(sentryPkgQualPath, "WithTransactionSource").Call(Qual(sentryPkgQualPath, "SourceURL")),
		),
		Id("transaction").Op(":=").Qual(sentryPkgQualPath, "StartTransaction").Call(
			Id("ctx"),
			Qual("fmt", "Sprintf").Call(Lit("%s %s"), Id("request").Dot("Method"), Id("request").Dot("URL").Dot("Path")),
			Id("options").Op("..."),
		),
		Defer().Id("transaction").Dot("Finish").Call(),
		Id("request").Op("=").Id("request").Dot("WithContext").Call(Id("transaction").Dot("Context").Call()),
		Id("hub").Dot("Scope").Call().Dot("SetRequest").Call(Id("request")),
		Defer().Func().Call().Block(panicDeferCallbackBlock...).Call(),
	}
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
	f.ImportName(sentryPkgQualPath, "sentry")
	if !noHeader {
		f.HeaderComment(fmt.Sprintf("Code generated by %s. DO NOT EDIT.", pkg))
	}
	f.Var().Defs(
		// Commentf("ErrAbort can be used to stop the middleware chain execution.").
		Line().
			Id("ErrAbort").Op("=").Qual("errors", "New").Call(Lit("chain execution stopped")),
	)
	for i := 1; i <= n; i++ {
		structName := chainStructName(i)
		f.
			// Commentf("%s provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [%s.Finally] by satisfying [net/http.HandlerFunc]", structName, structName).
			Line().
			Type().
			Id(structName).
			Types(genericTypes(i)...).
			Struct(fnParams(i)...)

		f.Line()

		// f.Commentf("ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing result%s of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.", lo.Ternary(i > 1, "s", ""))
		f.
			Func().
			Params(Id("chain").Id(structName).Types(parameterGenericTypes(i)...)).
			Id("ServeHTTP").
			Params(
				Id("response").Qual("net/http", "ResponseWriter"),
				Id("request").Add(Op("*")).Qual("net/http", "Request"),
			).
			Block(
				append(
					sentryOnRequest(sentryOnRequestPanicDeferCallbackBlock()),
					append(
						lo.Flatten(
							lo.Times(i-1, func(j int) []Code {
								return []Code{
									List(
										Id(genericTypeParamName(j)),
										Err(),
									).
										Op(":=").
										Id("chain").
										Dot(fnName(j + 1)).
										Call(
											append(
												[]Code{
													Id("response"),
													Id("request"),
													Id("hub"),
													Id("transaction"),
												},
												lo.Times(j, func(k int) Code { return Id(genericTypeParamName(k)) })...,
											)...,
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
						Id("_").
							Op("=").
							Id("chain").
							Dot(fnName(i)).
							Call(
								append(
									[]Code{
										Id("response"),
										Id("request"),
										Id("hub"),
										Id("transaction"),
									},
									lo.Times(i-1, func(j int) Code { return Id(genericTypeParamName(j)) })...,
								)...,
							),
					)...,
				)...,
			)

		f.Line()

		// f.Commentf("Finally executes middleware functions registered via [%s] in order, passing result%s of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.", factoryFuncName(i), lo.Ternary(i > 1, "s", ""))
		f.
			Func().
			Params(Id("chain").Id(structName).Types(parameterGenericTypes(i)...)).
			Id("Finally").
			Params(
				Id("catch").
					Func().
					Params(
						Qual("net/http", "ResponseWriter"),
						Add(Op("*")).Qual("net/http", "Request"),
						Add(Op("*")).Qual(sentryPkgQualPath, "Hub"),
						Add(Op("*")).Qual(sentryPkgQualPath, "Span"),
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
							append(
								sentryOnRequest(sentryOnRequestPanicDeferCallbackBlockFinally()),
								append(
									lo.Flatten(
										lo.Times(i-1, func(j int) []Code {
											return []Code{
												List(
													Id(genericTypeParamName(j)),
													Err(),
												).
													Op(":=").
													Id("chain").
													Dot(fnName(j + 1)).
													Call(
														append(
															[]Code{
																Id("response"),
																Id("request"),
																Id("hub"),
																Id("transaction"),
															},
															lo.Times(j, func(k int) Code { return Id(genericTypeParamName(k)) })...,
														)...,
													),
												If(
													Nil().
														Op("!=").
														Err(),
												).
													Block(
														If(
															Op("!").
																Add().
																Qual("errors", "Is").
																Call(Err(), Id("ErrAbort")),
														).
															Block(
																Id("catch").Call(Id("response"), Id("request"), Id("hub"), Id("transaction"), Err()),
															),
														Return(),
													),
											}
										}),
									),
									If(
										Err().
											Op(":=").
											Id("chain").
											Dot(fnName(i)).
											Call(
												append(
													[]Code{
														Id("response"),
														Id("request"),
														Id("hub"),
														Id("transaction"),
													},
													lo.Times(i-1, func(j int) Code { return Id(genericTypeParamName(j)) })...,
												)...,
											),
										Nil().Op("!=").Err(),
									).
										Block(
											If(
												Op("!").
													Add().
													Qual("errors", "Is").
													Call(Err(), Id("ErrAbort")),
											).
												Block(
													Id("catch").Call(Id("response"), Id("request"), Id("hub"), Id("transaction"), Err()),
												),
										),
								)...,
							)...,
						),
				),
			)

		f.Line()

		// f.Commentf("%s creates a chain of exactly %d function%s that will be executed in order.", factoryFuncName(i), i, lo.Ternary(i > 1, "s", ""))
		f.Func().
			Id(factoryFuncName(i)).
			Types(genericTypes(i)...).
			Params(fnParams(i)...).
			Id(structName).Types(parameterGenericTypes(i)...).
			Block(
				Return(
					Id(structName).
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
