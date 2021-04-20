/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   banking.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 11:11:26 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/20 11:40:14 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package banking

// If it starts with an upper case it's public, else it's private. Genius!

// Account struct
type Account struct {
	Owner   string
	Balance int
}
