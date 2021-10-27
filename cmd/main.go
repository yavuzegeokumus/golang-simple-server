package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Key struct {
	Key string `json:"key"`
}

var key Key

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func setKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&key)
		check(err)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "key set"}`))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "this method is not allowed"}`))
	}
}

func getKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		if key.Key == "" {
			w.WriteHeader(http.StatusNoContent)
		} else {
			message := `{"message": "key get successful", "data":"` + key.Key + `"}`

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(message))
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "this method is not allowed"}`))
	}
}

func flush(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if key.Key == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			key = Key{}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "key deleted"}`))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message": "this method is not allowed"}`))
	}
}

func saveKey(filename string, data string) {
	file, err := os.Create(filename)

	check(err)

	fmt.Println("File created " + filename)

	defer func() {
		err := file.Close()
		check(err)
	}()

	n, err := file.WriteString(data)
	check(err)
	fmt.Printf("wrote %d bytes\n", n)
}

func restoreKey() {
	path := "tmp/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir|0755)
	}

	files, err := ioutil.ReadDir(path)

	check(err)

	if len(files) > 0 {
		backupFileName := files[len(files)-1].Name()
		dat, err := os.ReadFile(path + backupFileName)
		check(err)

		key.Key = string(dat)
	}
}

func main() {
	restoreKey()

	http.HandleFunc("/setKey", setKey)
	http.HandleFunc("/getKey", getKey)
	http.HandleFunc("/flush", flush)

	ticker := time.NewTicker(60 * 60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if key.Key != "" {
					time := strconv.Itoa(int(time.Now().Unix()))
					filename := "tmp/" + time + "-data.txt"
					saveKey(filename, key.Key)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}
