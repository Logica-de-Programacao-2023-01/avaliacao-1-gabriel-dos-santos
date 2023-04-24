package q1

import (
	"errors"
	"fmt"
	"strconv"
)

//Em um dia quente de verão, Pete e seu amigo Billy decidiram comprar uma melancia. Eles escolheram a maior e mais
//saborosa, na opinião deles, e, em seguida, pesaram a fruta nas balanças, obtendo seu peso em quilos. Morrendo de sede,
//correram para casa com a melancia e decidiram dividi-la. No entanto, enfrentaram um problema difícil.
//
//Como grandes fãs de números pares, Pete e Billy querem dividir a melancia de tal forma que cada uma das duas partes pese
//um número par de quilos, sem que as partes necessariamente tenham pesos iguais. Mas os meninos estão extremamente
//cansados e querem começar a refeição o mais rápido possível. Portanto, precisam de ajuda para descobrir se é possível
//dividir a melancia da maneira que desejam. É importante destacar que cada um deles deve receber uma parte de peso
//positivo.
//
//A função deve retornar um valor booleano, indicando se é possível ou não dividir a melancia da forma desejada. Se o peso
//da melancia for menor ou igual a 0, a função deve retornar um erro.

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, errors.New("O peso da melancia dos amigos deve ser maior que zero")
	}
	if weight%2 == 0 && (weight/2)%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	weightString := "10"
	weight, err := strconv.Atoi(weightString)
	if err != nil {
		fmt.Println("Erro ao converter peso para int")
		return
	}

	divide, err := DivideWatermelon(weight)

	if err != nil {
		fmt.Println(err)
	} else if divide {
		fmt.Println("É possível dividir a melancia em duas partes pares")
	} else {
		fmt.Println("Não foi possível dividir em duas partes pares")
	}
}
