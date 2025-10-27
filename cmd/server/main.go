package main

import (
	"encoding/json"
	"net/http"
)



func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	response := "Hello!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}



func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", nil)
}