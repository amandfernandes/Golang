// 4. Crie dois programas, onde cada um define um novo tipo com base no tipo int. No primeiro programa, declare uma variável x utilizando a palavra-chave var e imprima tanto o valor quanto o tipo da variável no console usando fmt.Printf. No segundo programa, atribua o valor 42 à variável x e imprima-a no console utilizando fmt.Printf, destacando tanto o valor quanto o tipo da variável.

// Segundo Programa
package main

import "fmt"

type gun int

var x gun

func main() {
	x = 42
	fmt.Printf("x: %v\t%T", x, x)
}
