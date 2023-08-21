# Middle

Go HTTP middleware I like to use

## Usage

Checkout the [examples](./examples/) directory for usage examples.

## Compatibility

As this package only depends on [`net/http`](https://pkg.go.dev/net/http) package APIs, it can be used in any framework or libraries that uses types exposed by [`net/http`](https://pkg.go.dev/net/http) package, e.g., [`github.com/julienschmidt/httprouter`](https://github.com/julienschmidt/httprouter).

## Limitations

This package exposes middleware functions chain with up to 26 functions, i.e., `Ware1` to `Ware26`. I'm planning to expose do some cleanups and make the code generator ready for public use, so you generate your `middle.WareN` up to any number of `N` you need.
