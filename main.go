package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// fmt.Print("Enter query: ")
	// var q string
	// fmt.Scanln(&q)

	u := "https://jsonplaceholder.typicode.com/posts/"

	resp, err := http.Get(u)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(body))

	var td []User
	err = json.Unmarshal(body, &td)
	fmt.Printf("%v\n", td)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i, v := range td {
		fmt.Printf("Index: %v\n", td[i])
		fmt.Printf("UserId: %v\n", v.UserId)
		fmt.Printf("Id: %v\n", v.Id)
		fmt.Printf("Body: %v\n", v.Body)
		fmt.Printf("Title: %v\n", v.Title)
	}
}
