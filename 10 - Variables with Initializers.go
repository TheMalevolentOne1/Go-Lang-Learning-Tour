package VariablesWithInitialisers

import "fmt"

var i, j int = 1, 2 // Explict Type

func main() {
	var c, py, js = 1, true, "no!" // Implicit Type

	fmt.Println(i, j, c, py, js)
}
