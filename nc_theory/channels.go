/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   channels.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 21:20:07 by lpaulo-m         ###   ########.fr       */
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
	turboHotRodCrash()
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

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
