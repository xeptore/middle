package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeptore/middle"
)

func main() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/path", middle.Ware3(m1, m2, m3).Then(handler))
	http.ListenAndServe("127.0.0.1:1080", router)
}

func m1(http.ResponseWriter, *http.Request) (string, error) {
	fmt.Println("in m1!")
	return "First!", nil
}

func m2(http.ResponseWriter, *http.Request) (int, error) {
	fmt.Println("in m2!")
	return -42, nil
}

func m3(http.ResponseWriter, *http.Request) (bool, error) {
	fmt.Println("in m3!")
	return true, nil
}

func handler(w http.ResponseWriter, r *http.Request, s string, i int, b bool) {
	w.Write([]byte(fmt.Sprintf("%s\n%d\n%v", s, i, b)))
}
