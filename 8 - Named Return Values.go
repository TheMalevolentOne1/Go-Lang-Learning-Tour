package NamedReturnValues

import "fmt"

func add(n1, n2 int) (total int) {
	return n1 + n2
}

func main() {
	fmt.Println(add(5, 5))
}
