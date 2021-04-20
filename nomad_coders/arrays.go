/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   arrays.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 04:10:48 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run arrays.go
// go build arrays.go && ./arrays

package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Arrays ===")

	array_demo()
	fmt.Println("---")
	slice_demo()
}

func array_demo() {
	names := [8]string{"gar", "mon", "bo", "zia", "fire", "walk", "with", "me"}

	names[1] = "laura"
	names[3] = "palmer"
	names[5] = "dale"
	names[7] = "cooper"
	fmt.Println(names)
}

func slice_demo() {
	names := []string{"gar", "mon", "bo", "zia", "fire", "walk", "with", "me"}

	fmt.Println(names)
	names = append(names, "laura")
	names = append(names, "palmer")
	names = append(names, "dale")
	names = append(names, "cooper")
	fmt.Println(names)
}
