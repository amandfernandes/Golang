// 3. Desenvolva um programa que inicializa três variáveis (x, y, z) com valores específicos (42, “James Bond” e true, respectivamente). Em seguida, utiliza a função fmt.Sprintf para criar uma string formatada contendo informações sobre essas variáveis, incluindo seus valores e os nomes das variáveis. A string formatada é armazenada na variável s e, por fim, é impressa no console usando fmt.Println.

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("x: %v\ty: %v\tz: %v", x, y, z)
	fmt.Println(s)
}
