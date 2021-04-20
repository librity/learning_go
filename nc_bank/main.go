/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 11:11:28 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 11:46:21 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run main

package main

import (
	"fmt"

	"github.com/librity/learning_go/nc_bank/banking"
)

func main() {
	account := banking.Account{Owner: "Lior", Balance: 1000}
	fmt.Println(account)
}
