package main

import (
    "net/http"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"


	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

// main function program starts here
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Creating 2 entities to store the retrieved data and can use wherever necessary
type Receipt struct {
	ID           string  `json:"id"` // Id is added to use in GetPoints API
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        float64 `json:"total"`
	Items        []Item  `json:"items"`
	Points       int     `json:"points"` //points is added, to store calculatePoints value.
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

var receipts map[string]Receipt // Global variable to store receipts

func init() {
	receipts = make(map[string]Receipt)
}


func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	receipt.ID = uuid.New().String()
	receipt.Points = calculatePoints(receipt)
	receipts[receipt.ID] = receipt

	response := map[string]string{"id": receipt.ID}
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	receipt, found := receipts[id]
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": receipt.Points}
	json.NewEncoder(w).Encode(response)
}

// Calculate points based on rules mentioned in the exercise

func calculatePoints(receipt Receipt) int {
	points := len(receipt.Retailer)

	if isRoundDollar(receipt.Total) {
		points += 50
	}

	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	points += len(receipt.Items) / 2 * 5

	for _, item := range receipt.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLen%3 == 0 {
			points += int(item.Price*0.2 + 0.5)
		}
	}

	purchaseTime := strings.Split(receipt.PurchaseTime, ":")
	hour := purchaseTime[0]
	if hourInt := hourToInt(hour); hourInt >= 14 && hourInt < 16 {
		points += 10
	}

	purchaseDay := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
	if dayToInt(purchaseDay)%2 != 0 {
		points += 6
	}

	return points
}

//writing functions to find the nature of the amount in the receipt
func isRoundDollar(amount float64) bool {
	return amount == float64(int(amount))
}

func isMultipleOfQuarter(amount float64) bool {
	return math.Mod((amount*100), 25.0) == 0.0
}

func hourToInt(hour string) int {
	var h int
	fmt.Sscanf(hour, "%d", &h)
	return h
}

func dayToInt(day string) int {
	var d int
	fmt.Sscanf(day, "%d", &d)
	return d
}
