// Escreva uma função que receba um slice de inteiros e retorne um novo slice contendo
// apenas os valores únicos, ou seja, sem duplicatas. Utilize um mapa para auxiliar na
// remoção das duplicatas.
package main

import "fmt"

func remov(lista []int) []int {
	mapa := make(map[int]bool)
	for _, numero := range lista {
		mapa[numero] = true
	}
	slice2 := []int{}
	for add := range mapa {
		slice2 = append(slice2, add)
	}
	return slice2
}
func main() {

	lista := []int{1, 2, 3, 4, 4, 3, 2}
	a := remov(lista)
	fmt.Print(lista)
	fmt.Print(a)
}
