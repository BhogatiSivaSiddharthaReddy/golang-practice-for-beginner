package main

import "fmt"

// In real net/http package this becomes:
//
// http.Request
//
// Request contains:
// - headers
// - body
// - method
// - URL
//
// We are only keeping Path for learning.
type Request struct {
	Path string
}

// Equivalent to:
//
// http.ResponseWriter
//
// Real ResponseWriter writes data
// back to TCP connection/socket.
type Responsewriter struct{}

// Simulating writing response back to client.
func (rw Responsewriter) Write(data string) {
	fmt.Println(data)
}

// This is equivalent to:
//
//	type Handler interface {
//	    ServeHTTP(ResponseWriter, *Request)
//	}
//
// This interface is the foundation
// of Go's entire HTTP package.
//
// Anything implementing ServeHTTP()
// becomes a valid HTTP handler.
type Handler interface {
	ServeHTTP(Responsewriter, *Request)
}

// Struct-based handler.
//
// Similar to:
//
// type MyHandler struct{}
//
// in real net/http usage.
type structhandler struct{}

// Since structhandler has ServeHTTP(),
// it satisfies Handler interface.
//
// Equivalent to custom handlers in real Go.
func (sh structhandler) ServeHTTP(
	rw Responsewriter,
	r *Request,
) {

	rw.Write("hello World")
}

// This is equivalent to:
//
// type HandlerFunc func(ResponseWriter, *Request)
//
// inside real net/http package.
//
// This is the trick that allows
// normal functions to behave like handlers.
type Handlefunc func(Responsewriter, *Request)

// This method exists in real net/http too.
//
// Because of this method,
// Handlefunc satisfies Handler interface.
//
// So normal functions can be used
// wherever Handler is expected.
func (H Handlefunc) ServeHTTP(
	RW Responsewriter,
	R *Request,
) {

	// Execute original function.
	H(RW, R)
}

// Normal function handler.
//
// Equivalent to:
//
// func hello(
//
//	w http.ResponseWriter,
//	r *http.Request,
//
// )
//
// in real HTTP applications.
func hello(rw Responsewriter, r *Request) {
	rw.Write("Hello from function")
}

// Equivalent to:
//
// http.ServeMux
//
// Router stores:
//
// route -> handler mapping
type router struct {
	list map[string]Handler
}

// Similar to creating new ServeMux.
//
// Real package internally initializes
// routing structures too.
func create() *router {
	return &router{
		list: make(map[string]Handler),
	}
}

// Equivalent to:
//
// mux.Handle()
//
// Stores handler inside routing table.
func (r *router) handle(
	s string,
	h Handler,
) {

	r.list[s] = h
}

// Equivalent to:
//
// mux.HandleFunc()
//
// Converts normal function into
// Handlefunc type internally.
//
// This is exactly what real
// net/http package does.
func (r *router) handlefunc(
	s string,
	f func(r Responsewriter, re *Request),
) {

	// Function converted into Handlefunc
	// so it satisfies Handler interface.
	r.handle(s, Handlefunc(f))
}

// Equivalent to:
//
// func (mux *ServeMux) ServeHTTP(...)
//
// Router itself now behaves as handler.
//
// Real ServeMux also implements
// Handler interface.
func (r *router) ServeHTTP(
	w Responsewriter,
	re *Request,
) {

	// Route lookup.
	//
	// Real net/http checks request path
	// and finds matching handler.
	handler := r.list[re.Path]

	// Dynamic dispatch happens here.
	//
	// handler is interface type.
	//
	// Actual concrete type may be:
	// - structhandler
	// - Handlefunc
	//
	// Go calls correct ServeHTTP()
	// depending on concrete type stored
	// inside interface.
	handler.ServeHTTP(w, re)
}

func main() {
	//serverexample function call
	server_example()
	// Interface variable.
	//
	// Equivalent to:
	//
	// var h http.Handler
	var Hello Handler

	// Structhandler satisfies Handler
	// because it has ServeHTTP().
	Hello = structhandler{}

	req := Request{"/"}

	req_pointer := &req

	res := Responsewriter{}

	// Dynamic dispatch.
	//
	// Interface internally contains:
	//
	// type  -> structhandler
	// value -> structhandler{}
	//
	// So Go calls:
	//
	// structhandler.ServeHTTP()
	Hello.ServeHTTP(res, req_pointer)

	var hel Handler

	// hello function converted into
	// Handlefunc type.
	//
	// Equivalent to:
	//
	// http.HandlerFunc(hello)
	hel = Handlefunc(hello)

	// Dynamic dispatch again.
	//
	// Interface contains:
	//
	// type  -> Handlefunc
	// value -> hello function
	//
	// Go calls:
	//
	// Handlefunc.ServeHTTP()
	//
	// which internally executes:
	//
	// hello()
	hel.ServeHTTP(res, req_pointer)

	// Creating router.
	//
	// Equivalent to:
	//
	// http.NewServeMux()
	rout := create()

	// Registering route.
	//
	// Equivalent to:
	//
	// mux.HandleFunc("/hello", hello)
	rout.handlefunc("/hello", hello)

	// Manual route lookup.
	//
	// Real routers do this internally.
	temp := rout.list["/hello"]

	// Calling matched handler manually.
	temp.ServeHTTP(res, req_pointer)

	// Printing stored interface value.
	fmt.Println(rout.list["/hello"])

	// Router itself behaves as handler now.
	//
	// Equivalent to:
	//
	// mux.ServeHTTP(...)
	//
	// Real net/http server eventually
	// calls ServeHTTP() on router.
	rout.ServeHTTP(res, req_pointer)

	//serverexample function call
	server_example()
}
