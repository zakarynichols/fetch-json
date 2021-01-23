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

	var posts []User
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i, p := range posts {
		fmt.Printf("%v\n", posts[i].UserId)
		if p.UserId >= 7 {
			fmt.Printf("Id: %v\n", p.Id)
			fmt.Printf("Body: %v\n", p.Body)
			fmt.Printf("Title: %v\n", p.Title)
		}
	}
}
