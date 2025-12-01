package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func CustomerServer(request *http.Request, writer *httptest.ResponseRecorder) {
	fmt.Fprint(writer, "20000")
}
