package main

import "net/http"

func handlerReadlines(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
