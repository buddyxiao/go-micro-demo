package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type registryRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	waite := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waite.Add(1)
		go registry(registryRequest{
			Username: fmt.Sprintf("username%d", rand.Int63n(1000)),
			Password: "dslkjfdlks",
		})
	}
	waite.Wait()
}

func registry(s registryRequest) {
	marshal, err := json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/v1/user/registry", bytes.NewReader(marshal))
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{Timeout: time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	all, err := ioutil.ReadAll(response.Body)
	log.Println(bytes.NewBuffer(all).String())
}
