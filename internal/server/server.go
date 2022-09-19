package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/TovarischSuhov/go-synchronius-example/internal/api"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not allowed method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var msg api.Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Duration(msg.Sleep) * time.Second)
	resp := api.Response{Message: msg.Message}
	buf, err := json.Marshal(resp)
	w.Write(buf)
	w.WriteHeader(http.StatusOK)
}
