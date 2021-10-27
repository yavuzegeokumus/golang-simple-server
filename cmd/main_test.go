package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var domain = "http://localhost:9000"

func TestSetKey(t *testing.T) {
	want := "{\"message\": \"key set\"}"
	var endpoint = domain + "/setKey"
	var jsonStr = []byte(`{"key":"test"}`)

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	got := string(body)
	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestGetKey(t *testing.T) {
	want := "{\"message\": \"key get successful\", \"data\":\"test\"}"
	var endpoint = domain + "/getKey"

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	got := string(body)

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSaveKey(t *testing.T) {
	want := "test"
	filename := "/tmp/testfile.txt"

	saveKey(filename, want)

	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	got := string(file)

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestFLush(t *testing.T) {
	want := "{\"message\": \"key deleted\"}"
	var endpoint = domain + "/flush"

	req, err := http.NewRequest("DELETE", endpoint, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	got := string(body)

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
