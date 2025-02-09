package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println("start to handler req...")
	defer fmt.Println("handle end.")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
		interErr := http.StatusInternalServerError
		http.Error(w, err.Error(), interErr)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8091", nil)
}
