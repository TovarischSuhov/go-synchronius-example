package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TovarischSuhov/go-synchronius-example/internal/api"
)

var defaultClient http.Client

func SendMessage(messageName string, sleepTime int) {
	msg := api.Message{Message: messageName, Sleep: sleepTime}
	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := defaultClient.Post("http://localhost:8080/ping", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(r))
}
