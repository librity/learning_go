/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   channels.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/09/16 22:06:50 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run channerls.go
// go build channerls.go && ./channerls

package main

import (
	"fmt"
	"time"
)

func main() {
	hotOrNot()
	receiveAllMessages()
	deadlock()
}

func hotOrNot() {
	fmt.Println("\n=== Manually block for messages ===")

	// [...] declares a fixed-size array
	sexys := [...]string{"lior", "money"}
	channel := make(chan string)

	for _, sexy := range sexys {
		go isSexy(sexy, channel)
	}

	resultOne := <-channel
	fmt.Println("Waiting for messages")
	fmt.Println("Received this message:", resultOne)
	fmt.Println("Received this message:", <-channel)
}

func receiveAllMessages() {
	fmt.Println("\n=== Automatically block for messages ===")

	sexys := [...]string{"lior", "money", "tubers", "whiskey", "waffles", "stake"}
	channel := make(chan string)

	for _, sexy := range sexys {
		go isSexy(sexy, channel)
	}

	fmt.Println("Waiting for messages")
	for i := 0; i < len(sexys); i++ {
		fmt.Println("Received this message:", <-channel)
	}
}

func deadlock() {
	fmt.Println("\n=== Should crash here ===")

	sexys := [...]string{"lior", "money"}
	channel := make(chan string)

	for _, sexy := range sexys {
		go isSexy(sexy, channel)
	}

	result := <-channel
	fmt.Println(result)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isSexy(sexyThing string, channel chan string) {
	time.Sleep(time.Second * 2)
	channel <- sexyThing + " is very sexy!"
}
