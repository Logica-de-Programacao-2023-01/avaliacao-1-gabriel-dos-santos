package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

import "fmt"

func DominoPieces(M, N int) (int, error) {
	if M <= 0 || N <= 0 {
		return 0, fmt.Errorf("M and N must be greater than 0")
	}
	area := M * N
	if area%2 == 0 {
		return area / 2, nil
	} else {
		return (area - 1) / 2, nil
	}
}

func main() {
	max, err := DominoPieces(2, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(max)
	}
}
