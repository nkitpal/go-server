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
