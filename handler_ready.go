package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	respondeWithJSON(w, 200, struct{}{})
}