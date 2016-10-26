package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add", add).Methods("GET")
	router.HandleFunc("/sub", sub).Methods("GET")
	router.HandleFunc("/mul", mul).Methods("GET")
	router.HandleFunc("/div", div).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func add(w http.ResponseWriter, req *http.Request) {
	op1 := req.URL.Query().Get("A")
	op2 := req.URL.Query().Get("B")
	a, _ := strconv.ParseFloat(op1, 64)
	b, _ := strconv.ParseFloat(op2, 64)
	w.WriteHeader(200)
	w.Write([]byte(FloatToString(a + b)))
}
func sub(w http.ResponseWriter, req *http.Request) {
	op1 := req.URL.Query().Get("A")
	op2 := req.URL.Query().Get("B")
	a, _ := strconv.ParseFloat(op1, 64)
	b, _ := strconv.ParseFloat(op2, 64)
	w.WriteHeader(200)
	w.Write([]byte(FloatToString(a - b)))
}
func mul(w http.ResponseWriter, req *http.Request) {
	op1 := req.URL.Query().Get("A")
	op2 := req.URL.Query().Get("B")
	a, _ := strconv.ParseFloat(op1, 64)
	b, _ := strconv.ParseFloat(op2, 64)
	w.WriteHeader(200)
	w.Write([]byte(FloatToString(a * b)))
}
func div(w http.ResponseWriter, req *http.Request) {
	op1 := req.URL.Query().Get("A")
	op2 := req.URL.Query().Get("B")
	a, _ := strconv.ParseFloat(op1, 64)
	b, _ := strconv.ParseFloat(op2, 64)
	w.WriteHeader(200)
	w.Write([]byte(FloatToString(a / b)))
}

// to convert a float number to a string
func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}
