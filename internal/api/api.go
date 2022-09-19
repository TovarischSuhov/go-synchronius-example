package api

type Message struct {
	Sleep   int    `json:"sleep"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}
