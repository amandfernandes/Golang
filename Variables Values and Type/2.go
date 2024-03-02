// 2. Desenvolva um programa que declare e inicialize três variáveis (x, y e z) com tipos int, string e bool, respectivamente. Utilize a função fmt.Printf para imprimir no console os valores das variáveis, formatando a saída de modo a incluir os nomes das variáveis e seus valores. Observe como a saída do programa destaca a ausência de valor das variáveis, evidenciando qual conceito?

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("x: %v\ny: %v\nz: %v", x, y, z)
}

//Conceito Zero Value (Valor Zero)
