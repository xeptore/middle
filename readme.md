# Middle

Go HTTP middleware I like to use.

## Usage

Checkout the [examples](./examples/) directory for usage examples.

## Compatibility

As this package only depends on [`net/http`](https://pkg.go.dev/net/http) package APIs, it can be used in any framework or libraries that uses types exposed by [`net/http`](https://pkg.go.dev/net/http) package, e.g., [`github.com/julienschmidt/httprouter`](https://github.com/julienschmidt/httprouter).

## Limitations

This package exposes middleware functions chain builders for up to 26 functions, i.e., `Chain1` to `Chain26`. Although I think this is way more than enough for many, not most of, applications, you can generate your `middle.WareN` up to any number of `N` you need using [generator](#using-generator).

## Using Generator

See [`gen`](./gen/) package flags, and [`./gen.go`](./gen.go) for an example of usage.
