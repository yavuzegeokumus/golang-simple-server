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

	if r.Method == "POST" {

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
	}
}

func fileExists(filename string) bool {
	flag, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !flag.IsDir()
}

func saveKey(filename string) {
	file, err := os.Create(filename)

	check(err)

	fmt.Println("File created " + filename)

	defer func() {
		err := file.Close()
		check(err)
	}()

	n, err := file.WriteString(key.Key)
	fmt.Printf("wrote %d bytes\n", n)
}

func restoreKey() {
	files, err := ioutil.ReadDir("/tmp/")

	check(err)

	if len(files) > 0 {
		backupFileName := files[len(files)-1].Name()
		dat, err := os.ReadFile("/tmp/" + backupFileName)
		check(err)

		key.Key = string(dat)
	}
}

func main() {
	restoreKey()

	http.HandleFunc("/setKey", setKey)
	http.HandleFunc("/getKey", getKey)
	http.HandleFunc("/flush", flush)

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if key.Key != "" {
					time := strconv.Itoa(int(time.Now().Unix()))
					filename := "/tmp/" + time + "-data.txt"
					saveKey(filename)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	http.ListenAndServe(":8081", nil)
}
