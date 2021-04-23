/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 19:29:37 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 20:51:40 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run urls.go main.go

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

var errRequestFailed = errors.New("request failed")

func main() {
	results := hitUrls(demoUrls)
	printResults(results)

	results = hitUrls(stressUrls)
	printResults(results)
}

func printResults(results map[string]string) {
	color.Blue("\n=== Results ===\n")
	for url, result := range results {
		printUrl(url)
		if result == "OFFLINE" {
			printOffline()
			continue
		}
		printOnline()
	}
}

func hitUrls(urls []string) map[string]string {
	var results = make(map[string]string)

	color.Blue("\n=== Hitting Urls ===\n")
	for _, url := range urls {
		result := "ONLINE"
		err := hitUrl(url)
		if err != nil {
			result = "OFFLINE"
		}
		results[url] = result
	}

	return results
}

func hitUrl(url string) error {
	fmt.Println("Cheking:", url)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		// fmt.Println("ERROR:", err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
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
