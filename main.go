package main

import (
	"./rssfilter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	config, err := rssfilter.ParseConfig(data)
	if err != nil {
		log.Fatal(err)
	}

	server := &rssfilter.Server{
		Config: config,
	}
	http.Handle("/", server)
	http.ListenAndServe(":8080", nil)
}
