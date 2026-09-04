package MultiParameterShorthand

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("5 + 5 =", add(5, 5))
}
