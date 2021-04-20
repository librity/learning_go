/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 11:11:28 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 12:49:54 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

// go run main

package main

import (
	"fmt"
	"log"

	"github.com/librity/learning_go/nc_bank/accounts"
)

func main() {
	bitcoin_wallet_demo()
	account_demo()
}

func bitcoin_wallet_demo() {
	fmt.Println("=== BitcoinWallet ===")

	bitcoin_wallet := accounts.BitcoinWallet{Owner: "Lior", Balance: 0}

	bitcoin_wallet.Owner = "Nigerian Prince"
	bitcoin_wallet.Balance = 1000000
	fmt.Println(bitcoin_wallet)
}

func account_demo() {
	fmt.Println("=== Secure Account ===")

	secure_account := accounts.NewAccount("Lior")
	fmt.Println(secure_account)

	secure_account.ChangeOwner("Pepe")
	fmt.Printf("New owner: %s\n", secure_account.Owner())

	secure_account.Deposit(10)
	fmt.Printf("New balance: %d\n", secure_account.Balance())

	secure_account.Withdraw(5)
	fmt.Println("New balance: ", secure_account.Balance())

	fmt.Println(secure_account)

	error := secure_account.Withdraw(100)
	if error != nil {
		log.Fatalln(error)
	}
	fmt.Printf("New balance: %d\n", secure_account.Balance())
}
