package main

import (
	"io"
	"net/http"
)

type StringHandler struct{
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request){
		io.WriteString(writer, sh.message)
		Printfln("Method: %v", request.Method)
		Printfln("URL: %v", request.URL)
		Printfln("HTTP Version: %v", request.Proto)
		Printfln("Host: %v", request.Host)
		for name, val := range request.Header{
			Printfln("Header: %v, Value: %v", name, val)
		}
		Printfln("---")
	}

func main(){
	// for _, p := range Products{
	// 	Printfln("Product: %v, Category: %v, Price: $%.2f", 
	// 	p.Name, p.Category, p.Price)
	// }

	err := http.ListenAndServe(":5000", StringHandler{message: "Hello Elijah"})
	if err != nil{
		Printfln("Error: %v", err.Error())
	}
}