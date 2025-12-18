package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
func main(){
	go startServer() // run server in background
	go postServer()

	fmt.Println("web server started")
	PerformGetRequest()
	PerformPostRequest()

}

func postServer() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintln(w, "POST request received")
	})

	fmt.Println("POST server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func startServer() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server")
	})

	http.ListenAndServe(":8080", nil)
}


// handling get request
func PerformGetRequest(){
	const myurl ="http://localhost:8080/get"
	response , err :=http.Get(myurl)

	if(err!= nil){
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)

	content , err := ioutil.ReadAll(response.Body)
	if(err!= nil){
		fmt.Println(err)
		panic(err)
	}	
	fmt.Println("Response body: ", string(content))
}

//handling post request
func PerformPostRequest(){
	const myurl = "http://localhost:8080/post"

	// fake json payload
	requestBody := strings.NewReader(`{
		"name": "Ayush",
		"course": "Btech"
	}`)
	response , err :=http.Post(myurl , "application/json" , requestBody)
	if(err!= nil){
		fmt.Println(err)
		panic(err)
	}			
	
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)

}