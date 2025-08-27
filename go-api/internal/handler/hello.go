package handler

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
