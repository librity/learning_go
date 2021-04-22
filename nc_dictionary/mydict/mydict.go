/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   mydict.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 12:52:18 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 03:04:25 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("word not found")
var errWordExists = errors.New("word already exists")

// Search definition of a word
func (d Dictionary) Search(word string) (string, error) {
	value, wordExists := d[word]
	if wordExists {
		return value, nil
	}
	return "", errNotFound
}

// Add a definition of a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		return errWordExists
	}
	d[word] = def
	return nil
}
