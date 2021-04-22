/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   mydict.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/20 12:52:18 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 16:56:21 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("word not found")
	errWordExists = errors.New("word already exists")
	errCantUpdate = errors.New("can't update an undefined word")
	errCantDelete = errors.New("can't delete an undefined word")
)

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
	if wordExists(err) {
		return errWordExists
	}
	d[word] = def
	return nil
}

// Update the definition of a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if !wordExists(err) {
		return errCantUpdate
	}
	d[word] = def
	return nil
}

// Delete the definition of a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if !wordExists(err) {
		return errCantDelete
	}
	delete(d, word)
	return nil
}

func wordExists(err error) bool {
	return err == nil
}
