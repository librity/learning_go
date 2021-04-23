/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strings.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 21:03:37 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run strings.go
// go build strings.go && ./strings

package main

import (
	"fmt"
	"strings"
)

var works bool = true

func main() {
	strings_demo()
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

	_, uppercase := len_and_upper(first_name)
	fmt.Println(uppercase)

	length, uppercase := special_landu(last_name)
	fmt.Println(length, uppercase)

	repeate_me("gar", "mon", "bo", "zia")
}

func len_and_upper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func special_landu(name string) (length int, uppercase string) {
	defer fmt.Println("DEBUG: special_landu just executed")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeate_me(words ...string) {
	fmt.Println(words)
}
