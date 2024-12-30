package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responsewithError(w, 400, "This is an error message that is returned to the client")
}
