/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   structs.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 10:54:12 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 11:10:45 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run structs.go
// go build structs.go && ./structs

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Structs ===")

	structs_demo()
}

type Person struct {
	name           string
	age            int
	favorite_foods []string
}

func structs_demo() {
	booze := []string{"jack", "jameson", "kirsch", "smirnoff"}
	lior := Person{"lior", 42, booze}

	fmt.Println(lior)
	fmt.Println(lior.name)
	fmt.Println(lior.age)
	fmt.Println(lior.favorite_foods)
	fmt.Println("---")

	mass_media := []string{
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia", "garmonbozia",
		"garmonbozia", "garmonbozia", "garmonbozia"}
	lior_after_rehab := Person{name: "lior", age: 42, favorite_foods: mass_media}

	fmt.Println(lior_after_rehab)
	fmt.Println(lior_after_rehab.name)
	fmt.Println(lior_after_rehab.age)
	fmt.Println(lior_after_rehab.favorite_foods)
	fmt.Println("---")
}
