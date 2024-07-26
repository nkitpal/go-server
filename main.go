package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("making get req..")
	PerformGet()
	fmt.Println("making post req..")
	PerformPost()
}

func PerformGet() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code:", response.StatusCode)
	var resBuilder strings.Builder //hard way using string Builders

	content, _ := io.ReadAll(response.Body) // easy way less efficient
	fmt.Println(string(content))
	resBuilder.Write(content)

	fmt.Println(resBuilder.String())

}

func PerformPost() { // json data
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`{
		"coursename" : "GO with go",
		"duration" : "10 days"
		}	
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseBuilder strings.Builder
	content, _ := io.ReadAll(response.Body)
	responseBuilder.Write(content)

	fmt.Println(responseBuilder.String())

}
