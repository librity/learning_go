/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   Main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 01:57:13 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 02:09:48 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run Main.go
// go build Main.go && ./Main

package main

import "fmt"

var up_here int = 10

func main() {
	fmt.Println("Hello GO!")
	numbers()
	strings()
}

func numbers() {
	var a int
	a = 41
	var b int = 42
	c := 43

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", up_here, up_here)

	var j float32 = 27
	k := 99.
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
}

func strings() {

}
