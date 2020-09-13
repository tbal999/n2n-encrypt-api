package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type stringcount struct {
	String string `json:"String"`
	Key    []int  `json:"Key"`
	Output []byte `json:"Result"`
}

func getPort() string {
	p := os.Getenv("PORT")
	fmt.Println(p)
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func randomNumber(min, max int) int {
	z := seededRand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

func check(n *[]int, y int) {
	nest := *n
	insertion := randomNumber(1, y)
	for i := range nest {
		if nest[i] == insertion {
			return
		}
	}
	nest = append(nest, insertion)
	*n = nest
}

func encode(w http.ResponseWriter, r *http.Request) {
	S := stringcount{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
	err2 := json.Unmarshal(body, &S)
	if err2 != nil {
		fmt.Println(err2)
	}
	By := []byte(S.String)
	samelength := false
	for samelength == false {
		if len(By) == len(S.Key) {
			samelength = true
		}
		check(&S.Key, len(By)+1)
	}
	for index := range By {
		for index2 := range S.Key {
			S.Output = append(S.Output, byte(int(By[index])+S.Key[index2]))
		}
	}
	S.String = "Encoded"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(S)
}

func decode(w http.ResponseWriter, r *http.Request) {
	S := stringcount{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	err2 := json.Unmarshal(body, &S)
	if err2 != nil {
		fmt.Println(err2)
	}
	decoded := []byte{}
	counter := 0
	for i := range S.Output {
		counter = counter + 1
		if counter == len(S.Key) {
			decoded = append(decoded, byte(int(S.Output[i])-S.Key[counter-1]))
			counter = 0
		}
	}
	S.String = string(decoded[:])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(S)
}

func front(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Encodes and decodes strings!\n")
	fmt.Fprintf(w, "Commands:\n")
	fmt.Fprintf(w, `1) /encode/{string} - encodes a string and generates a key that's the length of the string, and the encoded output in bytes - example: {"String":"TEST","Key":[1,2,3,4],"Result":"encryptedoutputinbytes"}`+"\n")
	fmt.Fprintf(w, `2) /decode/JSON - decodes JSON formatted like above`+"\n")

}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var allroutes = routes{
	route{
		"encode",
		"POST",
		"/encode",
		encode,
	},
	route{
		"decode",
		"POST",
		"/decode",
		decode,
	},
}

func main() {
	//port := getPort()
	fmt.Println("API has started.")
	//fmt.Println("Running on port " + port)
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range allroutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.HandleFunc("/", front)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router)) //testing kraken git push
	//another test to check account
}
