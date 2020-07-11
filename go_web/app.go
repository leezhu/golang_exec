package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/julienschmidt/httprouter"
)

//MyHandler myhandler
// type MyHandler struct{}

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world")
// }

// type worldHandler struct{}

// func (wh *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "my world")
// }
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// fmt.Fprintf(w, "hello world,%s", p.ByName("hello"))
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	// fmt.Fprintln(w, r.Header)
	// fmt.Fprintln(w, string(body))
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

//类似于python的装饰器，
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(name)
		h(w, r)
		// h()
	}
}

//类似于python的装饰器，
func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(name)
		h(w, r)
		// h()
	}
}

func main() {

	mux := httprouter.New()
	server := http.Server{
		Addr:    "127.0.0.1:8001",
		Handler: mux,
	}
	mux.POST("/w/:hello", world)
	// handler := &MyHandler{}
	// worldHandler := &worldHandler{}
	// http.Handle("/h", handler)
	// http.Handle("/w", worldHandler)

	//handlerfunc 方式
	// http.HandleFunc("/w", hello)
	// http.HandleFunc("/h", world)

	// //串联技术
	// http.HandleFunc("/l", protect(log(hello)))

	//路由服务
	server.ListenAndServe()
}
