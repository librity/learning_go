/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   pointers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 04:02:00 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run pointers.go
// go build pointers.go && ./pointers

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Pointers ===")

	address()
	fmt.Println("---")
	pointers_demo()
}

func address() {
	a := 2
	b := a
	fmt.Println(a, b)

	a = 10
	fmt.Println(a, b)

	fmt.Println(&a, &b)
}

func pointers_demo() {
	a := 2
	b := &a
	var c *int = b
	fmt.Println(&a, &b, &c)
	fmt.Println(a, b, c)
	fmt.Println(&a, b, c)
	fmt.Println(a, *b, *c)

	a = 10
	fmt.Println(a, *b, *c)

	*b = 42
	fmt.Println(a, *b, *c)

	*c = 100
	fmt.Println(a, *b, *c)
}
