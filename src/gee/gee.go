package gee

/*
Package gee provides a lightweight web framework for building HTTP servers.

Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
Package net/http provides HTTP client and server implementations.
Package log implements a simple logging package.
*/
import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request);

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc // router map key: method + path, value: handler function 
}

// New is the constructor of gee.Engine
func New() * Engine{
	return &Engine{router: make(map[string]HandlerFunc)} // return a pointer to Engine, and initialize the router
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler;
}

// GET defines the method to add GET requests
func (engine *Engine) GET(pattern string, handler HandlerFunc){
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST requests
 func (engine *Engine) POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler);
 }

// PUT defines the method to add PUT requests
func (engine *Engine) PUT(pattern string, handler HandlerFunc){
	engine.addRoute("PUT", pattern, handler);
}

// DELETE defines the method to add DELETE requests
func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path;
	if handler, ok := engine.router[key]; ok {
		handler(w,req);
	}else{
		w.WriteHeader(http.StatusNotFound);
		_, _ = w.Write([]byte("404 NOT FOUND: " + req.URL.Path + "\n"))
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL);
	}
}
