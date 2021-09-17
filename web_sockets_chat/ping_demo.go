package main

import (
	"fmt"
	"net/http"
)

func pingDemo(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	PanicError(err)

	for {
		fmt.Println("Awaiting message...")
		_, payload, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\"\n", payload)
		fmt.Println("---")
	}
}
