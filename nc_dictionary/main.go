/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 12:52:03 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 03:04:16 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run main.go

package main

import (
	"fmt"

	"github.com/librity/learning_go/nc_dictionary/mydict"
)

func main() {
	typeDemo()
	searchDemo()
	addDemo()
}

func typeDemo() {
	fmt.Println("=== Map string ===")
	dictionary := mydict.Dictionary{}

	dictionary["hello"] = "world"
	fmt.Println(dictionary)
}

func searchDemo() {
	fmt.Println("=== Types have methods too ===")

	dictionary := mydict.Dictionary{"ham": "cheese"}

	definition, err := dictionary.Search("ham")
	handleError(definition, err)

	definition, err = dictionary.Search("pork")
	handleError(definition, err)
}

func addDemo() {
	fmt.Println("=== Add words to dictionary ===")

	dictionary := mydict.Dictionary{}

	dictionary.Add("pork", "beans")
	definition, err := dictionary.Search("pork")
	handleError(definition, err)

	err = dictionary.Add("pork", "spam")
	handleError("", err)
	definition, err = dictionary.Search("pork")
	handleError(definition, err)
}

func handleError(result string, err error) {
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("ERROR:", err)
	}
}
