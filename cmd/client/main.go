package main

import (
	"fmt"

	"github.com/TovarischSuhov/go-synchronius-example/internal/client"
)

func main() {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Example%d", i)
		sleep := i % 3
		client.SendMessage(name, sleep)
	}
}
