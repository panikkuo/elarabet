package main

import (
	"github.com/gorilla/mux"
	handler "github.com/panikkuo/elarabet/back-core/src/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Handler)
}
