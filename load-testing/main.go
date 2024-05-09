// add_service.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Assume we receive a JSON payload like {"a": 2, "b": 3}
	var data map[string]int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := data["a"] + data["b"]
	jsonResponse := map[string]int{"result": result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	// Assume we receive a JSON payload like {"a": 5, "b": 4}
	log.Print("chegooou !")
	var data map[string]int
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := data["a"] * data["b"]
	jsonResponse := map[string]int{"result": result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	fmt.Println("Add service running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
