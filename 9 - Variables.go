package Variables

import "fmt"

var num1, num2, num3 int // Package-Level Declaration

func main() {
	var i int // Function-Scope Declaration
	fmt.Println(i, num1, num2, num3)
}
