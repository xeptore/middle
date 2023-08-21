package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeptore/middle/v4"
)

func main() {
	router := httprouter.New()
	router.Handler("GET", "/path-that-ignores-handler-error", middle.Chain4(m1, m2, m3, handler))
	router.Handler("GET", "/path-that-handles-handler-error", middle.Chain4(m1, m2, m3, handler).Finally(unexpectedErrHandle))
	http.ListenAndServe("127.0.0.1:1080", router)
}

func m1(http.ResponseWriter, *http.Request) (string, error) {
	fmt.Println("in m1!")
	return "First!", nil
}

func m2(http.ResponseWriter, *http.Request, string) (int, error) {
	fmt.Println("in m2!")
	return -42, nil
}

func m3(http.ResponseWriter, *http.Request, string, int) (bool, error) {
	fmt.Println("in m3!")
	return true, nil
}

func handler(w http.ResponseWriter, r *http.Request, s string, i int, b bool) error {
	w.Write([]byte(fmt.Sprintf("%s\n%d\n%v", s, i, b)))
	return nil
}

func unexpectedErrHandle(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unexpected error: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal error!"))
}
