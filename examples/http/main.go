package main

import (
	"log"
	"net/http"

	"github.com/getsentry/sentry-go"
)

func main() {
	mux := http.NewServeMux()
	sentry.Init(sentry.ClientOptions{})
	hub := sentry.CurrentHub()
	Chain1Sentry(hub, handler)
	mux.Handle("/path-that-ignores-handler-error", nil)
	// mux.Handle("/path-that-handles-handler-error", middle.Chain4(m1, m2, m3, handler).Finally(unexpectedErrHandle))
	http.ListenAndServe("127.0.0.1:1080", mux)
}

// func m1(http.ResponseWriter, *http.Request) (string, error) {
// 	fmt.Println("in m1!")
// 	return "First!", nil
// }

// func m2(http.ResponseWriter, *http.Request, string) (int, error) {
// 	fmt.Println("in m2!")
// 	return -42, nil
// }

// func m3(http.ResponseWriter, *http.Request, string, int) (bool, error) {
// 	fmt.Println("in m3!")
// 	return true, nil
// }

func handler(w http.ResponseWriter, r *http.Request, hub *sentry.Hub) error {
	hub.CaptureMessage("fuck!")
	// w.Write([]byte(fmt.Sprintf("%s\n%d\n%v", s, i, b)))
	return nil
}

func unexpectedErrHandle(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unexpected error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal error!"))
}
