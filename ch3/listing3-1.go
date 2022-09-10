package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	//defer r1.Body.Close()
	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()

	/*
		r2, err := http.Head("http://www.google.com/robots.txt")
		//read responce body not shown
		defer r2.Body.Close()

		form := url.Values{}
		form.Add("foo", "bar")
		r3, err := http.Post(
			"https://www.google.com/robots.txt",
			"application/x-www-form-urlencoded",
			strings.NewReader(form.Encode()),
		)
		//read responce body not shown
		defer r3.Body.Close()
	*/
}
