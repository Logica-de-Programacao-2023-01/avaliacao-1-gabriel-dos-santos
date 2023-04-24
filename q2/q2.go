package q2

//Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de
//programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os
//amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso
//contrário, os amigos não escreveriam a solução do problema.
//
//Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você
//receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e
//Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false"
//indica o contrário.
//
//Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.

import "fmt"

func ProblemsSolved(answers [][3]bool) int {
	numSolved := 0
	for _, row := range answers {
		numTrue := 0
		for _, answer := range row {
			if answer {
				numTrue++
			}
		}
		if numTrue >= 2 {
			numSolved++
		}
	}
	return numSolved
}

func main() {
	answers := [][3]bool{{true, false, true}, {true, true, false}, {false, true, true}, {true, true, true}}
	numProblemsSolved := ProblemsSolved(answers)
	fmt.Printf("Os amigos resolveram %d problemas", numProblemsSolved)
}
