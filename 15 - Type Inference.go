package TypeInferences

import "fmt"

func main() {
	i := 120
	f := 12.5
	c := 0.867 + 0.5i

	fmt.Printf("Integer: %T \n", i) // int
	fmt.Printf("Float: %T \n", f)   // float64
	fmt.Printf("Complex: %T \n", c) // complex128
}
