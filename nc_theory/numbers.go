/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   numbers.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 02:36:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 03:39:17 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run numbers.go
// go build numbers.go && ./numbers

package main

import (
	"fmt"
)

var precise float64 = 42.42

func main() {
	numbers_demo()
}

func numbers_demo() {
	fmt.Println("=== NUMBERS ===")

	fmt.Println(multiply_f64(precise, precise))

	fmt.Println(multiply(2, int(precise)))
	fmt.Println(multiply(2, 3))

	total := roided_add(1, 2, 3, 4, 5, 6, 7, 88)
	fmt.Println(total)

	total = add(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)

	fmt.Println(can_i_drink(20))
	fmt.Println(can_i_drink(21))

	fmt.Println(can_i_drink_in_korea(15))
	fmt.Println(can_i_drink_in_korea(16))

	fmt.Println(can_i_drink_in_korea_v2(15))
	fmt.Println(can_i_drink_in_korea_v2(16))

	fmt.Println(can_i_drink_in_italy(9))
	fmt.Println(can_i_drink_in_italy(10))
}

func multiply(a, b int) int {
	return (a * b)
}

func multiply_f64(a, b float64) float64 {
	return (a * b)
}

func roided_add(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}

func add(numbers ...int) (total int) {
	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}
	return
}

func can_i_drink(age int) bool {
	if age < 21 {
		return false
	}
	return true
}

func can_i_drink_in_korea(age int) bool {
	if korean_age := age + 2; korean_age < 18 {
		return false
	}
	return true
}

func can_i_drink_in_korea_v2(age int) bool {
	switch korean_age := age + 2; {
	case korean_age < 18:
		return false
	case korean_age >= 18:
		return true
	}
	return false
}

func can_i_drink_in_italy(age int) bool {
	switch {
	case age < 10:
		return false
	case age == 10:
		return true
	case age > 10:
		return true
	}
	return false
}
