# Middle

Go HTTP middleware I like to use.

## Usage

Checkout the [examples](./examples/) directory for usage examples.

## Compatibility

As this package only depends on [`net/http`](https://pkg.go.dev/net/http) package APIs, it can be used in any framework or libraries that uses types exposed by [`net/http`](https://pkg.go.dev/net/http) package, e.g., [`github.com/julienschmidt/httprouter`](https://github.com/julienschmidt/httprouter).

## Limitations

This package exposes middleware functions chain builders for up to 27 functions, i.e., `Chain1` up to `Chain27`. Although I think this is way more than enough for most of applications, I plan to improve the generator so you can generate your `middle.ChainN` up to any number of `N` you need by using it. See [generator](#using-generator) for more.

## Using Generator

See [`gen`](./gen/gen.go) command line flags, and [`./gen.go`](./gen.go) for an example of usage.
