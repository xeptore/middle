package main

import (
	"net/http"
)

type HubScoper interface {
	SetRequest(*http.Request)
}

type ClonedHub[T HubScoper] interface {
	Scope() T
	// Hub[T]
}

type Hub[X HubScoper, Z ClonedHub[X], T any] interface {
	Clone() Z
	// Scope() HubScoper
}

type ChainHandler1[X HubScoper, Z ClonedHub[X], H Hub[X, Z, H]] struct {
	hub H
	f1  func(http.ResponseWriter, *http.Request, Z) error
}

func (chain ChainHandler1[X, Z, H]) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	hub := chain.hub.Clone()
	// options := []sentry.SpanOption{
	// 	sentry.WithOpName("http.server"),
	// 	sentry.ContinueFromRequest(r),
	// 	sentry.WithTransactionSource(sentry.SourceURL),
	// }

	// transaction := sentry.StartTransaction(ctx,
	// 	fmt.Sprintf("%s %s", r.Method, r.URL.Path),
	// 	options...,
	// )
	// defer transaction.Finish()
	// r = r.WithContext(transaction.Context())
	hub.Scope().SetRequest(request)
	// 	defer func() {
	// 		// TODO(tracing): if the next handler.ServeHTTP panics, store
	// 		// information on the transaction accordingly (status, tag,
	// 		// level?, ...).
	// 	}()
	_ = chain.f1(response, request, hub)
}

// Finally executes middleware functions registered via [Chain1] in order, passing result(s) of all previous function calls to it. If any of the functions in the chain returns a non-nil error, the execution stops, and executes catch with that error. If the error is [ErrAbort] according to [errors.Is] semantics, it is ignored, and catch will not be called, although the chain execution stops.
// func (chain ChainHandler1[S]) Finally(catch func(http.ResponseWriter, *http.Request, error)) http.HandlerFunc {
// 	return func(response http.ResponseWriter, request *http.Request) {
// 		if err := chain.f1(response, request); nil != err {
// 			if !errors.Is(err, os.ErrPermission) {
// 				catch(response, request, err)
// 			}
// 		}
// 	}
// }

func Chain1Sentry[X HubScoper, Z ClonedHub[X], H Hub[X, Z, H]](hub H, f1 func(http.ResponseWriter, *http.Request, Z) error) ChainHandler1[X, Z, H] {
	return ChainHandler1[X, Z, H]{}
}
