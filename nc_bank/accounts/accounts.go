/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   accounts.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 11:11:26 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 12:01:47 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package accounts

import "errors"

// BitcoinWallet not secure at all
type BitcoinWallet struct {
	Owner   string
	Balance int
}

type account struct {
	owner   string
	balance int
}

var errBadWithdraw = errors.New("amount exceedes balance and credit limit")

// NewAccount creates a private account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Balance gets account balance
func (a account) Balance() int {
	return a.balance
}

// Deposit adds amount to balance
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw substract amount from balance
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errBadWithdraw
	}
	a.balance -= amount
	return nil
}
