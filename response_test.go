package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {

	go StartServer()
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
