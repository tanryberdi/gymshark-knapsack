package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gymshark-knapsack/pkg"

	"github.com/gorilla/mux"
)

var items []int

func main() {
	router := mux.NewRouter()

	// Define a handler function for logging requests
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the request body
			requestBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			// Log the request method, URI, and body
			log.Printf("%s %s - Body: %s", r.Method, r.RequestURI, string(requestBody))

			// Replace the request body with a new buffer
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}

	// Use the logging middleware for all routes
	router.Use(loggingMiddleware)

	router.HandleFunc("/api/items", getItems).Methods("GET")
	router.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/api/items", createItems).Methods("POST")

	// Calculate the packages for the customer
	router.HandleFunc("/api/calculate/{order}", getOrder).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}

// nolint: errcheck
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for index, item := range items {
		if index == id {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	order, err := strconv.Atoi(params["order"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//itemsToCustomer, packsToCustomer, usedPacks := pkg.Knapsack(items, order)
	/*
		To further illustrate the rules above, please consider this custom pack size example:
		items:[23, 31, 53]
		Items order: 263
		Correct Number of packs: 2x23, 7x31
		Incorrect answer: 5x53
	*/
	_, packsToCustomer, _ := pkg.Knapsack(items, order)

	_ = json.NewEncoder(w).Encode(packsToCustomer)
}

func createItems(w http.ResponseWriter, r *http.Request) {
	// new item struct
	type newItemsCollection struct {
		Items string `json:"items"`
	}
	w.Header().Set("Content-Type", "application/json")

	var newItems newItemsCollection
	err := json.NewDecoder(r.Body).Decode(&newItems)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	itemsAsString := strings.Split(newItems.Items, ",")
	items = nil
	for _, item := range itemsAsString {
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		items = append(items, itemInt)
	}

	_ = json.NewEncoder(w).Encode(newItems)
}
