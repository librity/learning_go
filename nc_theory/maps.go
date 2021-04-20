/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   maps.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 04:19:23 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run maps.go
// go build maps.go && ./maps

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== maps ===")

	map_demo()
}

func map_demo() {
	dude := map[string]string{"name": "dude", "age": "53"}
	fmt.Println(dude)

	prisoner := map[string]int{"name": 90, "age": 87}
	fmt.Println(prisoner)

	raw := map[float64]int{3.2312451: 90, .23231231321312: 87}
	fmt.Println(raw)

	for key, value := range raw {
		fmt.Printf("KEY %f\n", key)
		fmt.Printf("VAL %d\n", value)
	}

	for key, value := range dude {
		fmt.Printf("KEY %f\n", key)
		fmt.Printf("VAL %d\n", value)
	}
}
