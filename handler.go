package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Only POST requests allowed")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var body requestBody
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON format: %v", err)
		return
	}

	
	err = UpdateYaml(body.Subdomain)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error updating yaml: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Router updated for %s", body.Subdomain)
}
