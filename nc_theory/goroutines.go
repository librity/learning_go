/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   goroutines.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 21:21:48 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run goroutines.go
// go build goroutines.go && ./goroutines

package main

import (
	"fmt"
	"time"
)

func main() {
	// turtleSnail()
	// turboHotRod()
	turboHotRodCrash()
}

func turtleSnail() {
	sexyCount("lior")
	sexyCount("money")
	sexyCount("tubers")
	sexyCount("whiskey")
	sexyCount("waffles")
	sexyCount("stake")
}

// go routines execute as long main runs
func turboHotRod() {
	go sexyCount("lior")
	go sexyCount("money")
	go sexyCount("tubers")
	go sexyCount("whiskey")
	go sexyCount("waffles")
	sexyCount("stake")
}

func turboHotRodCrash() {
	go sexyCount("lior")
	go sexyCount("money")
	go sexyCount("tubers")
	go sexyCount("whiskey")
	go sexyCount("waffles")
	go sexyCount("stake")
	time.Sleep(time.Second * 5)
}

func sexyCount(sexyThing string) {
	for i := 0; i < 10; i++ {
		fmt.Println(sexyThing, "is sexy", i)
		time.Sleep(time.Second)
	}
}
