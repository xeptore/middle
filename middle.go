package middle

import (
	"errors"
	"net/http"
)

var (
	// ErrAbort can be used to stop the middleware chain execution.
	ErrAbort = errors.New("chain execution stopped")
)

// ChainHandler1 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler1.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler1 struct {
	f1 func(http.ResponseWriter, *http.Request) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes the handler function, passing request, and response to it.
func (chain ChainHandler1) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	_ = chain.f1(response, request)
}

// Finally executes middleware function registered via [Chain1], passing request, and response to it.
func (chain ChainHandler1) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if err := chain.f1(response, request); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain1 creates a chain of exactly 1 function that will be executed in order.
func Chain1(f1 func(http.ResponseWriter, *http.Request) error) ChainHandler1 {
	return ChainHandler1{f1}
}

// ChainHandler2 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler2.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler2[A any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler2[A]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	_ = chain.f2(response, request, a)
}

// Finally executes middleware functions registered via [Chain2] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler2[A]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f2(response, request, a); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain2 creates a chain of exactly 2 functions that will be executed in order.
func Chain2[A any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) error) ChainHandler2[A] {
	return ChainHandler2[A]{f1, f2}
}

// ChainHandler3 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler3.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler3[A any, B any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler3[A, B]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	_ = chain.f3(response, request, a, b)
}

// Finally executes middleware functions registered via [Chain3] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler3[A, B]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f3(response, request, a, b); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain3 creates a chain of exactly 3 functions that will be executed in order.
func Chain3[A any, B any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) error) ChainHandler3[A, B] {
	return ChainHandler3[A, B]{f1, f2, f3}
}

// ChainHandler4 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler4.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler4[A any, B any, C any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler4[A, B, C]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	_ = chain.f4(response, request, a, b, c)
}

// Finally executes middleware functions registered via [Chain4] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler4[A, B, C]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f4(response, request, a, b, c); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain4 creates a chain of exactly 4 functions that will be executed in order.
func Chain4[A any, B any, C any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) error) ChainHandler4[A, B, C] {
	return ChainHandler4[A, B, C]{f1, f2, f3, f4}
}

// ChainHandler5 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler5.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler5[A any, B any, C any, D any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler5[A, B, C, D]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	_ = chain.f5(response, request, a, b, c, d)
}

// Finally executes middleware functions registered via [Chain5] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler5[A, B, C, D]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f5(response, request, a, b, c, d); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain5 creates a chain of exactly 5 functions that will be executed in order.
func Chain5[A any, B any, C any, D any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) error) ChainHandler5[A, B, C, D] {
	return ChainHandler5[A, B, C, D]{f1, f2, f3, f4, f5}
}

// ChainHandler6 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler6.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler6[A any, B any, C any, D any, E any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler6[A, B, C, D, E]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	_ = chain.f6(response, request, a, b, c, d, e)
}

// Finally executes middleware functions registered via [Chain6] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler6[A, B, C, D, E]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f6(response, request, a, b, c, d, e); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain6 creates a chain of exactly 6 functions that will be executed in order.
func Chain6[A any, B any, C any, D any, E any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) error) ChainHandler6[A, B, C, D, E] {
	return ChainHandler6[A, B, C, D, E]{f1, f2, f3, f4, f5, f6}
}

// ChainHandler7 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler7.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler7[A any, B any, C any, D any, E any, F any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler7[A, B, C, D, E, F]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	_ = chain.f7(response, request, a, b, c, d, e, f)
}

// Finally executes middleware functions registered via [Chain7] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler7[A, B, C, D, E, F]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f7(response, request, a, b, c, d, e, f); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain7 creates a chain of exactly 7 functions that will be executed in order.
func Chain7[A any, B any, C any, D any, E any, F any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) error) ChainHandler7[A, B, C, D, E, F] {
	return ChainHandler7[A, B, C, D, E, F]{f1, f2, f3, f4, f5, f6, f7}
}

// ChainHandler8 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler8.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler8[A any, B any, C any, D any, E any, F any, G any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler8[A, B, C, D, E, F, G]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	_ = chain.f8(response, request, a, b, c, d, e, f, g)
}

// Finally executes middleware functions registered via [Chain8] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler8[A, B, C, D, E, F, G]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f8(response, request, a, b, c, d, e, f, g); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain8 creates a chain of exactly 8 functions that will be executed in order.
func Chain8[A any, B any, C any, D any, E any, F any, G any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) error) ChainHandler8[A, B, C, D, E, F, G] {
	return ChainHandler8[A, B, C, D, E, F, G]{f1, f2, f3, f4, f5, f6, f7, f8}
}

// ChainHandler9 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler9.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler9[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler9[A, B, C, D, E, F, G, H]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	_ = chain.f9(response, request, a, b, c, d, e, f, g, h)
}

// Finally executes middleware functions registered via [Chain9] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler9[A, B, C, D, E, F, G, H]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f9(response, request, a, b, c, d, e, f, g, h); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain9 creates a chain of exactly 9 functions that will be executed in order.
func Chain9[A any, B any, C any, D any, E any, F any, G any, H any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) error) ChainHandler9[A, B, C, D, E, F, G, H] {
	return ChainHandler9[A, B, C, D, E, F, G, H]{f1, f2, f3, f4, f5, f6, f7, f8, f9}
}

// ChainHandler10 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler10.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler10[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler10[A, B, C, D, E, F, G, H, I]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	_ = chain.f10(response, request, a, b, c, d, e, f, g, h, i)
}

// Finally executes middleware functions registered via [Chain10] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler10[A, B, C, D, E, F, G, H, I]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f10(response, request, a, b, c, d, e, f, g, h, i); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain10 creates a chain of exactly 10 functions that will be executed in order.
func Chain10[A any, B any, C any, D any, E any, F any, G any, H any, I any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) error) ChainHandler10[A, B, C, D, E, F, G, H, I] {
	return ChainHandler10[A, B, C, D, E, F, G, H, I]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10}
}

// ChainHandler11 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler11.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler11[A, B, C, D, E, F, G, H, I, J]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	_ = chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
}

// Finally executes middleware functions registered via [Chain11] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler11[A, B, C, D, E, F, G, H, I, J]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain11 creates a chain of exactly 11 functions that will be executed in order.
func Chain11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) error) ChainHandler11[A, B, C, D, E, F, G, H, I, J] {
	return ChainHandler11[A, B, C, D, E, F, G, H, I, J]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11}
}

// ChainHandler12 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler12.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler12[A, B, C, D, E, F, G, H, I, J, K]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	_ = chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
}

// Finally executes middleware functions registered via [Chain12] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler12[A, B, C, D, E, F, G, H, I, J, K]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain12 creates a chain of exactly 12 functions that will be executed in order.
func Chain12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) error) ChainHandler12[A, B, C, D, E, F, G, H, I, J, K] {
	return ChainHandler12[A, B, C, D, E, F, G, H, I, J, K]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12}
}

// ChainHandler13 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler13.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler13[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler13[A, B, C, D, E, F, G, H, I, J, K, L]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	_ = chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
}

// Finally executes middleware functions registered via [Chain13] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler13[A, B, C, D, E, F, G, H, I, J, K, L]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain13 creates a chain of exactly 13 functions that will be executed in order.
func Chain13[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) error) ChainHandler13[A, B, C, D, E, F, G, H, I, J, K, L] {
	return ChainHandler13[A, B, C, D, E, F, G, H, I, J, K, L]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13}
}

// ChainHandler14 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler14.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler14[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler14[A, B, C, D, E, F, G, H, I, J, K, L, M]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	_ = chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
}

// Finally executes middleware functions registered via [Chain14] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler14[A, B, C, D, E, F, G, H, I, J, K, L, M]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain14 creates a chain of exactly 14 functions that will be executed in order.
func Chain14[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) error) ChainHandler14[A, B, C, D, E, F, G, H, I, J, K, L, M] {
	return ChainHandler14[A, B, C, D, E, F, G, H, I, J, K, L, M]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14}
}

// ChainHandler15 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler15.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler15[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	_ = chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
}

// Finally executes middleware functions registered via [Chain15] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain15 creates a chain of exactly 15 functions that will be executed in order.
func Chain15[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) error) ChainHandler15[A, B, C, D, E, F, G, H, I, J, K, L, M, N] {
	return ChainHandler15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}
}

// ChainHandler16 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler16.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler16[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	_ = chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}

// Finally executes middleware functions registered via [Chain16] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain16 creates a chain of exactly 16 functions that will be executed in order.
func Chain16[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) error) ChainHandler16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O] {
	return ChainHandler16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16}
}

// ChainHandler17 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler17.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler17[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	_ = chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
}

// Finally executes middleware functions registered via [Chain17] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain17 creates a chain of exactly 17 functions that will be executed in order.
func Chain17[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) error) ChainHandler17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P] {
	return ChainHandler17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17}
}

// ChainHandler18 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler18.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler18[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	_ = chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
}

// Finally executes middleware functions registered via [Chain18] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain18 creates a chain of exactly 18 functions that will be executed in order.
func Chain18[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) error) ChainHandler18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q] {
	return ChainHandler18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18}
}

// ChainHandler19 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler19.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler19[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	_ = chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
}

// Finally executes middleware functions registered via [Chain19] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain19 creates a chain of exactly 19 functions that will be executed in order.
func Chain19[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) error) ChainHandler19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R] {
	return ChainHandler19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19}
}

// ChainHandler20 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler20.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler20[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	_ = chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
}

// Finally executes middleware functions registered via [Chain20] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain20 creates a chain of exactly 20 functions that will be executed in order.
func Chain20[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) error) ChainHandler20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S] {
	return ChainHandler20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20}
}

// ChainHandler21 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler21.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler21[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	_ = chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
}

// Finally executes middleware functions registered via [Chain21] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain21 creates a chain of exactly 21 functions that will be executed in order.
func Chain21[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) error) ChainHandler21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T] {
	return ChainHandler21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21}
}

// ChainHandler22 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler22.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler22[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	_ = chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
}

// Finally executes middleware functions registered via [Chain22] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain22 creates a chain of exactly 22 functions that will be executed in order.
func Chain22[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) error) ChainHandler22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U] {
	return ChainHandler22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22}
}

// ChainHandler23 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler23.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler23[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error)
	f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
	if nil != err {
		return
	}
	_ = chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
}

// Finally executes middleware functions registered via [Chain23] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain23 creates a chain of exactly 23 functions that will be executed in order.
func Chain23[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) error) ChainHandler23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V] {
	return ChainHandler23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23}
}

// ChainHandler24 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler24.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler24[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error)
	f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error)
	f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
	if nil != err {
		return
	}
	w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
	if nil != err {
		return
	}
	_ = chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
}

// Finally executes middleware functions registered via [Chain24] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain24 creates a chain of exactly 24 functions that will be executed in order.
func Chain24[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) error) ChainHandler24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W] {
	return ChainHandler24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24}
}

// ChainHandler25 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler25.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler25[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error)
	f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error)
	f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error)
	f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
	if nil != err {
		return
	}
	w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
	if nil != err {
		return
	}
	x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
	if nil != err {
		return
	}
	_ = chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x)
}

// Finally executes middleware functions registered via [Chain25] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain25 creates a chain of exactly 25 functions that will be executed in order.
func Chain25[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error), f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) error) ChainHandler25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X] {
	return ChainHandler25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25}
}

// ChainHandler26 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler26.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler26[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error)
	f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error)
	f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error)
	f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) (Y, error)
	f26 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
	if nil != err {
		return
	}
	w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
	if nil != err {
		return
	}
	x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
	if nil != err {
		return
	}
	y, err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x)
	if nil != err {
		return
	}
	_ = chain.f26(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y)
}

// Finally executes middleware functions registered via [Chain26] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		y, err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f26(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain26 creates a chain of exactly 26 functions that will be executed in order.
func Chain26[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error), f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) (Y, error), f26 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y) error) ChainHandler26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y] {
	return ChainHandler26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26}
}

// ChainHandler27 provides capability of processing chain functions in order by satisfying [net/http.Handler], or with an optional chain error handler via [ChainHandler27.Finally] by satisfying [net/http.HandlerFunc]
type ChainHandler27[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any, Z any] struct {
	f1  func(http.ResponseWriter, *http.Request) (A, error)
	f2  func(http.ResponseWriter, *http.Request, A) (B, error)
	f3  func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4  func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5  func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6  func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error)
	f9  func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error)
	f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error)
	f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error)
	f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error)
	f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error)
	f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error)
	f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error)
	f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error)
	f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error)
	f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error)
	f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error)
	f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error)
	f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error)
	f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error)
	f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error)
	f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error)
	f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) (Y, error)
	f26 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y) (Z, error)
	f27 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z) error
}

// ServeHTTP satisfies [net/http.Handler]. It executes functions in the chain in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops.
func (chain ChainHandler27[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	b, err := chain.f2(response, request, a)
	if nil != err {
		return
	}
	c, err := chain.f3(response, request, a, b)
	if nil != err {
		return
	}
	d, err := chain.f4(response, request, a, b, c)
	if nil != err {
		return
	}
	e, err := chain.f5(response, request, a, b, c, d)
	if nil != err {
		return
	}
	f, err := chain.f6(response, request, a, b, c, d, e)
	if nil != err {
		return
	}
	g, err := chain.f7(response, request, a, b, c, d, e, f)
	if nil != err {
		return
	}
	h, err := chain.f8(response, request, a, b, c, d, e, f, g)
	if nil != err {
		return
	}
	i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
	if nil != err {
		return
	}
	j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
	if nil != err {
		return
	}
	k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
	if nil != err {
		return
	}
	l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
	if nil != err {
		return
	}
	m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
	if nil != err {
		return
	}
	n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
	if nil != err {
		return
	}
	o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	if nil != err {
		return
	}
	p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	if nil != err {
		return
	}
	q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	if nil != err {
		return
	}
	r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	if nil != err {
		return
	}
	s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	if nil != err {
		return
	}
	t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	if nil != err {
		return
	}
	u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
	if nil != err {
		return
	}
	v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
	if nil != err {
		return
	}
	w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
	if nil != err {
		return
	}
	x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
	if nil != err {
		return
	}
	y, err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x)
	if nil != err {
		return
	}
	z, err := chain.f26(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y)
	if nil != err {
		return
	}
	_ = chain.f27(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z)
}

// Finally executes middleware functions registered via [Chain27] in order, passing results of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
func (chain ChainHandler27[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		c, err := chain.f3(response, request, a, b)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		d, err := chain.f4(response, request, a, b, c)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		e, err := chain.f5(response, request, a, b, c, d)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		f, err := chain.f6(response, request, a, b, c, d, e)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		g, err := chain.f7(response, request, a, b, c, d, e, f)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		h, err := chain.f8(response, request, a, b, c, d, e, f, g)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		i, err := chain.f9(response, request, a, b, c, d, e, f, g, h)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		j, err := chain.f10(response, request, a, b, c, d, e, f, g, h, i)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		k, err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		l, err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		m, err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		n, err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		o, err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		p, err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		q, err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		r, err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		s, err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		t, err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		u, err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		v, err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		w, err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		x, err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		y, err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		z, err := chain.f26(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y)
		if nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
			return
		}
		if err := chain.f27(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z); nil != err {
			if !errors.Is(err, ErrAbort) {
				catch(response, request, err)
			}
		}
	}
}

// Chain27 creates a chain of exactly 27 functions that will be executed in order.
func Chain27[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any, Z any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error), f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) (Y, error), f26 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y) (Z, error), f27 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z) error) ChainHandler27[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z] {
	return ChainHandler27[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27}
}
