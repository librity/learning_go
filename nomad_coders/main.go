/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 03:11:04 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run main.go
// go build main.go && ./main

package main

import (
	"fmt"
	"strings"
)

var works bool = true

func main() {
	strings_demo()
	numbers_demo()
}

func multiply(a, b int) int {
	return (a * b)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func minilandu(name string) (length int, uppercase string) {
	defer fmt.Println("DEBUG: minilandu just executed")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeateMe(words ...string) {
	fmt.Println(words)
}

func strings_demo() {
	fmt.Println("=== STRINGS ===")

	const first_name = "lior"
	const last_name string = "geniole"
	name := "Cooper"

	fmt.Println("Sup.")
	fmt.Printf("%s, %s\n", last_name, first_name)
	fmt.Println(name)
	fmt.Println(works)

	_, uppercase := lenAndUpper(first_name)
	fmt.Println(uppercase)

	length, uppercase := minilandu(last_name)
	fmt.Println(length, uppercase)

	repeateMe("gar", "mon", "bo", "zia")
}

func numbers_demo() {
	fmt.Println("=== NUMBERS ===")

	fmt.Println(multiply(2, 3))
}
