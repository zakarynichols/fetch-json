package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	UserId    string `json:"userId"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// p, _ := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", string(body))
}
