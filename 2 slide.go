// Escreva uma função que receba um mapa onde as chaves são nomes de alunos e os
// valores são slices de notas. A função deve calcular a média das notas de cada aluno e
// retornar um novo mapa onde as chaves são os nomes dos alunos e os valores são as
// médias correspondentes.
package main

func media(entrada map[string][]float64) map[string]float64 {
	me := make(map[string]float64)

	for aluno, notas := range entrada {
		var t float64 = 0
		for _, nota := range notas {
			t += nota
		}
		me[aluno] = t / float64(len(notas))
	}
	return me
}
func main() {

}
