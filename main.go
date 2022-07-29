package main

import (
	"fmt"
	"io"
	"net/http"
)

type stringHandler struct{
	message string
}
func (sh stringHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	io.WriteString(writer,sh.message)
}
func main() {
	err := http.ListenAndServe(":1234",stringHandler{message:"Hello World"})
	if err != nil{
		fmt.Println(err)
	}
}
