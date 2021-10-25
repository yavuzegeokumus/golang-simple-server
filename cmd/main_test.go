package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestSetKey(t *testing.T) {
	want := "{\"message\": \"key set\"}"

	var jsonStr = []byte(`{"key":"test"}`)
	req, err := http.NewRequest("POST", "http://localhost:8081/setKey", bytes.NewBuffer(jsonStr))
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

	req, err := http.NewRequest("GET", "http://localhost:8081/getKey", nil)

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
	filename := "/testoutput/testfile.txt"
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

	req, err := http.NewRequest("DELETE", "http://localhost:8081/flush", nil)

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
