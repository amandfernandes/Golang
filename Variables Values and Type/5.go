// 5. Desenvolva um programa que define um novo tipo com base no tipo int. O programa declara duas variáveis, x do tipo declarado e y do tipo int. Em seguida, o programa imprime no console o valor e o tipo da variável x. Posteriormente, atribui o valor 42 à variável x e imprime novamente seu valor e tipo. Em seguida, converte o valor da variável x para o tipo int, atribuindo-o à variável y, e imprime no console o valor e o tipo da variável y.

package main

import "fmt"

type gun int

var x gun
var y int

func main() {
	fmt.Printf("x: %v\n%T\n", x, x)
	x = 42
	fmt.Printf("x: %v\t%T\n", x, x)
	y = int(x)
	fmt.Printf("y: %v\t%T", y, y)
}
