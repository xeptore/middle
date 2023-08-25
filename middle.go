package middle

import "net/http"

type chain1 struct {
	f1 func(http.ResponseWriter, *http.Request) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain1) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	_ = chain.f1(response, request)
}

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain1) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if err := chain.f1(response, request); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain1 creates a chain of exactly 1 number of function that will be executed in order.
func Chain1(f1 func(http.ResponseWriter, *http.Request) error) chain1 {
	return chain1{f1}
}

type chain2[A any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain2[A]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	a, err := chain.f1(response, request)
	if nil != err {
		return
	}
	_ = chain.f2(response, request, a)
}

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain2[A]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			return
		}
		if err := chain.f2(response, request, a); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain2 creates a chain of exactly 2 number of functions that will be executed in order.
func Chain2[A any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) error) chain2[A] {
	return chain2[A]{f1, f2}
}

type chain3[A any, B any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain3[A, B]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain3[A, B]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		a, err := chain.f1(response, request)
		if nil != err {
			return
		}
		b, err := chain.f2(response, request, a)
		if nil != err {
			return
		}
		if err := chain.f3(response, request, a, b); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain3 creates a chain of exactly 3 number of functions that will be executed in order.
func Chain3[A any, B any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) error) chain3[A, B] {
	return chain3[A, B]{f1, f2, f3}
}

type chain4[A any, B any, C any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain4[A, B, C]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain4[A, B, C]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f4(response, request, a, b, c); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain4 creates a chain of exactly 4 number of functions that will be executed in order.
func Chain4[A any, B any, C any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) error) chain4[A, B, C] {
	return chain4[A, B, C]{f1, f2, f3, f4}
}

type chain5[A any, B any, C any, D any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain5[A, B, C, D]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain5[A, B, C, D]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f5(response, request, a, b, c, d); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain5 creates a chain of exactly 5 number of functions that will be executed in order.
func Chain5[A any, B any, C any, D any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) error) chain5[A, B, C, D] {
	return chain5[A, B, C, D]{f1, f2, f3, f4, f5}
}

type chain6[A any, B any, C any, D any, E any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain6[A, B, C, D, E]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain6[A, B, C, D, E]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f6(response, request, a, b, c, d, e); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain6 creates a chain of exactly 6 number of functions that will be executed in order.
func Chain6[A any, B any, C any, D any, E any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) error) chain6[A, B, C, D, E] {
	return chain6[A, B, C, D, E]{f1, f2, f3, f4, f5, f6}
}

type chain7[A any, B any, C any, D any, E any, F any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain7[A, B, C, D, E, F]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain7[A, B, C, D, E, F]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f7(response, request, a, b, c, d, e, f); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain7 creates a chain of exactly 7 number of functions that will be executed in order.
func Chain7[A any, B any, C any, D any, E any, F any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) error) chain7[A, B, C, D, E, F] {
	return chain7[A, B, C, D, E, F]{f1, f2, f3, f4, f5, f6, f7}
}

type chain8[A any, B any, C any, D any, E any, F any, G any] struct {
	f1 func(http.ResponseWriter, *http.Request) (A, error)
	f2 func(http.ResponseWriter, *http.Request, A) (B, error)
	f3 func(http.ResponseWriter, *http.Request, A, B) (C, error)
	f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error)
	f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error)
	f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error)
	f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error)
	f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) error
}

// ServeHTTP satisfies [net/http.Handler].
func (chain chain8[A, B, C, D, E, F, G]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain8[A, B, C, D, E, F, G]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f8(response, request, a, b, c, d, e, f, g); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain8 creates a chain of exactly 8 number of functions that will be executed in order.
func Chain8[A any, B any, C any, D any, E any, F any, G any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) error) chain8[A, B, C, D, E, F, G] {
	return chain8[A, B, C, D, E, F, G]{f1, f2, f3, f4, f5, f6, f7, f8}
}

type chain9[A any, B any, C any, D any, E any, F any, G any, H any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain9[A, B, C, D, E, F, G, H]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain9[A, B, C, D, E, F, G, H]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f9(response, request, a, b, c, d, e, f, g, h); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain9 creates a chain of exactly 9 number of functions that will be executed in order.
func Chain9[A any, B any, C any, D any, E any, F any, G any, H any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) error) chain9[A, B, C, D, E, F, G, H] {
	return chain9[A, B, C, D, E, F, G, H]{f1, f2, f3, f4, f5, f6, f7, f8, f9}
}

type chain10[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain10[A, B, C, D, E, F, G, H, I]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain10[A, B, C, D, E, F, G, H, I]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f10(response, request, a, b, c, d, e, f, g, h, i); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain10 creates a chain of exactly 10 number of functions that will be executed in order.
func Chain10[A any, B any, C any, D any, E any, F any, G any, H any, I any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) error) chain10[A, B, C, D, E, F, G, H, I] {
	return chain10[A, B, C, D, E, F, G, H, I]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10}
}

type chain11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain11[A, B, C, D, E, F, G, H, I, J]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain11[A, B, C, D, E, F, G, H, I, J]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f11(response, request, a, b, c, d, e, f, g, h, i, j); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain11 creates a chain of exactly 11 number of functions that will be executed in order.
func Chain11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) error) chain11[A, B, C, D, E, F, G, H, I, J] {
	return chain11[A, B, C, D, E, F, G, H, I, J]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11}
}

type chain12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain12[A, B, C, D, E, F, G, H, I, J, K]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain12[A, B, C, D, E, F, G, H, I, J, K]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f12(response, request, a, b, c, d, e, f, g, h, i, j, k); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain12 creates a chain of exactly 12 number of functions that will be executed in order.
func Chain12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) error) chain12[A, B, C, D, E, F, G, H, I, J, K] {
	return chain12[A, B, C, D, E, F, G, H, I, J, K]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12}
}

type chain13[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain13[A, B, C, D, E, F, G, H, I, J, K, L]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain13[A, B, C, D, E, F, G, H, I, J, K, L]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f13(response, request, a, b, c, d, e, f, g, h, i, j, k, l); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain13 creates a chain of exactly 13 number of functions that will be executed in order.
func Chain13[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) error) chain13[A, B, C, D, E, F, G, H, I, J, K, L] {
	return chain13[A, B, C, D, E, F, G, H, I, J, K, L]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13}
}

type chain14[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain14[A, B, C, D, E, F, G, H, I, J, K, L, M]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain14[A, B, C, D, E, F, G, H, I, J, K, L, M]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f14(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain14 creates a chain of exactly 14 number of functions that will be executed in order.
func Chain14[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) error) chain14[A, B, C, D, E, F, G, H, I, J, K, L, M] {
	return chain14[A, B, C, D, E, F, G, H, I, J, K, L, M]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14}
}

type chain15[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f15(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain15 creates a chain of exactly 15 number of functions that will be executed in order.
func Chain15[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) error) chain15[A, B, C, D, E, F, G, H, I, J, K, L, M, N] {
	return chain15[A, B, C, D, E, F, G, H, I, J, K, L, M, N]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15}
}

type chain16[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f16(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain16 creates a chain of exactly 16 number of functions that will be executed in order.
func Chain16[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) error) chain16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O] {
	return chain16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16}
}

type chain17[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f17(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain17 creates a chain of exactly 17 number of functions that will be executed in order.
func Chain17[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) error) chain17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P] {
	return chain17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17}
}

type chain18[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f18(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain18 creates a chain of exactly 18 number of functions that will be executed in order.
func Chain18[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) error) chain18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q] {
	return chain18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18}
}

type chain19[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f19(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain19 creates a chain of exactly 19 number of functions that will be executed in order.
func Chain19[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) error) chain19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R] {
	return chain19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19}
}

type chain20[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f20(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain20 creates a chain of exactly 20 number of functions that will be executed in order.
func Chain20[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) error) chain20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S] {
	return chain20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20}
}

type chain21[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f21(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain21 creates a chain of exactly 21 number of functions that will be executed in order.
func Chain21[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) error) chain21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T] {
	return chain21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21}
}

type chain22[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f22(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain22 creates a chain of exactly 22 number of functions that will be executed in order.
func Chain22[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) error) chain22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U] {
	return chain22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22}
}

type chain23[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f23(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain23 creates a chain of exactly 23 number of functions that will be executed in order.
func Chain23[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) error) chain23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V] {
	return chain23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23}
}

type chain24[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f24(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain24 creates a chain of exactly 24 number of functions that will be executed in order.
func Chain24[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) error) chain24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W] {
	return chain24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24}
}

type chain25[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f25(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain25 creates a chain of exactly 25 number of functions that will be executed in order.
func Chain25[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error), f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) error) chain25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X] {
	return chain25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25}
}

type chain26[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any] struct {
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

// ServeHTTP satisfies [net/http.Handler].
func (chain chain26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
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

// Finally executes handler registered via [Then] similar to [Then], and executes handle only if returned error from handler is not nil.
func (chain chain26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
		if err := chain.f26(response, request, a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y); nil != err {
			catch(response, request, err)
		}
	}
}

// Chain26 creates a chain of exactly 26 number of functions that will be executed in order.
func Chain26[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any](f1 func(http.ResponseWriter, *http.Request) (A, error), f2 func(http.ResponseWriter, *http.Request, A) (B, error), f3 func(http.ResponseWriter, *http.Request, A, B) (C, error), f4 func(http.ResponseWriter, *http.Request, A, B, C) (D, error), f5 func(http.ResponseWriter, *http.Request, A, B, C, D) (E, error), f6 func(http.ResponseWriter, *http.Request, A, B, C, D, E) (F, error), f7 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F) (G, error), f8 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G) (H, error), f9 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H) (I, error), f10 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I) (J, error), f11 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J) (K, error), f12 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K) (L, error), f13 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L) (M, error), f14 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M) (N, error), f15 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N) (O, error), f16 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) (P, error), f17 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) (Q, error), f18 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) (R, error), f19 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) (S, error), f20 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S) (T, error), f21 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T) (U, error), f22 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U) (V, error), f23 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V) (W, error), f24 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W) (X, error), f25 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X) (Y, error), f26 func(http.ResponseWriter, *http.Request, A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y) error) chain26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y] {
	return chain26[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y]{f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26}
}
