package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/TovarischSuhov/go-synchronius-example/internal/api"
	"github.com/TovarischSuhov/log"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Warn("Not allowed method")
		w.WriteHeader(http.StatusNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn(err)
	}
	var msg api.Message
	err := json.Unmarshal(body, &msg)
	if err != nil {
		log.Warn(err)
	}
	time.Sleep(msg.Sleep * time.Second)
	resp := api.Response{Message: msg.Message}
	buf, err := json.Marshal(resp)
	w.Write(buf)
	w.WriteHeader(http.StatusOK)
}
