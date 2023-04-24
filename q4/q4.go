package q4

import (
	"errors"
	"fmt"
)

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	num := len(prices)
	if num == 0 {
		return 0, errors.New("a lista está em branco")
	}
	if num == 1 {
		return 3, nil
	}
	increasing := true
	decreasing := true
	for i := 1; i < num; i++ {
		if prices[i] > prices[i-1] {
			decreasing = false
		} else if prices[i] < prices[i-1] {
			increasing = false
		}
		if !increasing && !decreasing {
			return 3, nil
		}
	}
	if increasing {
		return 1, nil
	}
	return 2, nil
}

func main() {
	prices1 := []int{1, 2, 3, 4}
	prices2 := []int{4, 3, 2, 1}
	prices3 := []int{1, 3, 2, 4}
	prices4 := []int{1}
	prices5 := []int{}
	fmt.Println(ClassifyPrices(prices1))
	fmt.Println(ClassifyPrices(prices2))
	fmt.Println(ClassifyPrices(prices3))
	fmt.Println(ClassifyPrices(prices4))
	fmt.Println(ClassifyPrices(prices5))
}
