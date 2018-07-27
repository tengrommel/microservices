package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	/*
	The HandleFunc method creates a Handler type on the DefaultServeMux handler,
	mapping the path passed in the first parameter to the function in the second parameter:
		func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	 */
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello, World\n")
}
