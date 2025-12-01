package main

import (
	"fmt"
	"net/http"
)

func CustomerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20000")
}
