// Escreva uma função que receba uma string como entrada e retorne um mapa onde as
// chaves são os caracteres presentes na string e os valores são o número de ocorrências
// de cada caractere.
package main

import (
	"fmt"
)

func ret(entrada string) map[string]int {
	ac := make(map[string]int)

	for _, l := range entrada {
		ac[string(l)]++
	}
	return ac
}
func main() {
	entrada := "Banana"
	resultado := ret(entrada)
	fmt.Print(resultado)
}
