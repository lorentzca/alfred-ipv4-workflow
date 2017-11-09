package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title string `json:"title"`
}

func main() {
	var ipv4 *string = flag.String("ipv4", "127.1.2.3/8", "ipv4 address")
	flag.Parse()

	_, ipv4Net, err := net.ParseCIDR(*ipv4)

	if err != nil {
		log.Fatal(err)
	}

	item := Item{Title: ipv4Net.String()}
	items := Items{Items: []Item{item}}

	jsonBytes, err := json.Marshal(items)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
