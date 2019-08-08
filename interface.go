package main

import "fmt"
import "net/http"

type counter int64

// implements http.Handler interface
func (ctr *counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr = *ctr + 1
	fmt.Fprintln(w, *ctr)
}

func main() {
	ctr := new(counter)
	http.Handle("/", ctr)
	http.ListenAndServe(":8080", nil)
}