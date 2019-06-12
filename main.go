package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

var (
	fAddr = flag.String("addr", ":8080", "Address to listen on")
)

func handler(rw http.ResponseWriter, r *http.Request) {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		rw.WriteHeader(500)
		fmt.Fprintln(rw, err.Error())
		return
	}
	fmt.Fprintln(rw, string(data))
}

func main() {
	if err := http.ListenAndServe(*fAddr, http.HandlerFunc(handler)); err != nil {
		panic(err)
	}
}
