package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getTokenId(w http.ResponseWriter, r *http.Request) {
	//Lets make an Http Post Request to jtb
	postBody, _ := json.Marshal(map[string]string{
		"email":      "info@totagotech.com",
		"password":   "Unlock*2021@",
		"clientname": "jtb",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://api.jtb.gov.ng:2089/api/GetTokenID", "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Println(sb)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func individualtinvalidation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenId := params["tokenId"]
	tin := params["tin"]

	postBody, _ := json.Marshal(map[string]string{
		"tin": tin,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://api.jtb.gov.ng:2089/api/individualtinvalidation?tokenid="+tokenId, "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// sb := string(body)
	// fmt.Println(sb)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func nonindividualtinvalidation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenId := params["tokenId"]
	tin := params["tin"]

	postBody, _ := json.Marshal(map[string]string{
		"tin": tin,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://api.jtb.gov.ng:2089/api/nonindividualtinvalidation?tokenid="+tokenId, "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// sb := string(body)
	// fmt.Println(sb)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/GetTokenID", getTokenId)
	router.HandleFunc("/api/individualtinvalidation/{tin}/{tokenId}", individualtinvalidation)
	router.HandleFunc("/api/nonindividualtinvalidation/{tin}/{tokenId}", nonindividualtinvalidation)

	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(corsObj)(router)))
}

func main() {
	fmt.Println("JTB Bridge Server Started")
	handleRequests()
}
