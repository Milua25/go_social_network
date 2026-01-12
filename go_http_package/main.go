package main

import (
	"fmt"
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Server Running!!")
	w.Write([]byte("Hello from the world of Go!"))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "404 not Found", 404)
		return
	}
	if req.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hola!!!")
}

func main() {

	s := &api{addr: ":8080"}

	//Initialize serverMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	mux.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting port at port 8080\n")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
