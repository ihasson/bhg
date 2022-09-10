package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Status struct {
	Message string
	Status  string
}

func main() {
	fmt.Println(len(os.Args), os.Args)
	ip := os.Args[1]
	port := os.Args[2]
	s := "http://" + ip + ":" + port + "/ping"
	fmt.Println(s)
	res, err := http.Post(
		"http://127.0.0.1/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
