/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 19:29:37 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 23:13:46 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run urls.go main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type HitResult struct {
	url    string
	online bool
}

func main() {
	hitAndPrint(demoUrls)
	hitAndPrint(stressUrls)
}

func hitAndPrint(urls []string) {
	c := make(chan HitResult)

	color.Blue("\n=== Hitting Urls ===\n")
	for _, url := range urls {
		fmt.Println("Hitting:", url)
		go hitUrl(c, url)
	}

	color.Blue("\n=== Results ===\n")
	for i := 0; i < len(urls); i++ {
		result := <-c
		printResult(result)
	}
}

// chan<- defines a send-only channel
func hitUrl(c chan<- HitResult, url string) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- HitResult{url: url, online: false}
		return
	}
	c <- HitResult{url: url, online: true}
}

func printResult(result HitResult) {
	printUrl(result.url)
	if result.online {
		printOnline()
		return
	}
	printOffline()
}

func printUrl(str string) {
	color.Set(color.FgCyan)
	fmt.Print(str)
	color.Unset()
}

func printOffline() {
	color.Set(color.FgRed)
	fmt.Print("	OFFLINE\n")
	color.Unset()
}

func printOnline() {
	color.Set(color.FgGreen)
	fmt.Print("	ONLINE\n")
	color.Unset()
}
