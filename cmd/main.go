package main

import (
	"encoding/json"
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

	router.HandleFunc("/api/items", getItems).Methods("GET")
	router.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/api/items", createItems).Methods("POST")

	// Calculate the packages for the customer
	router.HandleFunc("/api/calculate/{capacity}", getCapacity).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

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
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func getCapacity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	capacity, err := strconv.Atoi(params["capacity"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//itemsToCustomer, packsToCustomer, usedPacks := pkg.Knapsack(items, capacity)
	/*
		To further illustrate the rules above, please consider this custom pack size example:
		items:[23, 31, 53]
		Items order: 263
		Correct Number of packs: 2x23, 7x31
		Incorrect answer: 5x53
	*/
	_, packsToCustomer, _ := pkg.Knapsack(items, capacity)

	json.NewEncoder(w).Encode(packsToCustomer)
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

	itemsAsSttring := strings.Split(newItems.Items, ",")
	items = nil
	for _, item := range itemsAsSttring {
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		items = append(items, itemInt)
	}

	json.NewEncoder(w).Encode(newItems)
}
