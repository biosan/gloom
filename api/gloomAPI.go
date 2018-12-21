package main

import (
    "encoding/json"
    "log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/biosan/gloom"
)

// Item struct : A simple rapresentation of Add/Query requests body data
type Item struct {
    Data	string   `json:"data"`
}

// QueryResponse : Simple response for Query requests body
type QueryResponse struct {
	ProbablyIn	bool   `json:"probably_in"`
}

// InitParams : Parameters for BloomFilter initialization
type InitParams struct {
	K	uint	`json:"number_of_hash_functions"`
	M	uint	`json:"size_in_bits"`
}

var bf *gloom.BloomFilter

// QueryItemEndpoint : Send response '{probably_in : true}' if '{data: XXX}' is probably in the BF
// 					   Send response '{probably_in : false}' if '{data: XXX}' is definitely not in the BF
func QueryItemEndpoint(w http.ResponseWriter, req *http.Request) {
	// Decode request and query the Bloom Filter
	var item Item
	_ = json.NewDecoder(req.Body).Decode(&item)
	data := []byte(item.Data)
	query := bf.Query(data)

	// Setup and send response
	response := new(QueryResponse)
	response.ProbablyIn = query
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// AddItemEndpoint : Insert '{data: XXX}' inside the Bloom Filter
func AddItemEndpoint(w http.ResponseWriter, req *http.Request) {
	// Decode request and add to Bloom Filter
	var item Item
    _ = json.NewDecoder(req.Body).Decode(&item)
	bf.Add([]byte(item.Data))

	// Setup and send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
}

// CreateBloomFilterEndpoint : Create a new Bloom Filter, if there is still one, it will be destoryed
func CreateBloomFilterEndpoint(w http.ResponseWriter, req *http.Request) {
	// Decode request and create a new Bloom Filter
	var init InitParams
	_ = json.NewDecoder(req.Body).Decode(&init)
	bf = gloom.New(init.K, init.M)
	// Setup and send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}


func main() {
	// Create a base BloomFilter to serve requests even whitout initialize the BF using POST:/create
	// TODO : Delete this and make POST:/create mandatory by sending an
	//		  error when receiving add/query without an initialized BF.
	bf = gloom.New(3, 1024)
	// Create a new router and add endpoints
    router := mux.NewRouter()
    router.HandleFunc("/add", AddItemEndpoint).Methods("POST")
    router.HandleFunc("/query", QueryItemEndpoint).Methods("POST")
    router.HandleFunc("/create", CreateBloomFilterEndpoint).Methods("POST")
	//router.HandleFunc("/nuke", NukeBloomFilterEndpoint).Methods("DELETE")
	// TODO : Make 
    log.Fatal(http.ListenAndServe(":8888", router))
}