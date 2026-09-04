package ZeroValues

import "fmt"

var (
	boolZero  bool
	intZero   int
	floatZero float64
	strZero   string
)

func main() {
	fmt.Println("Zero Types:")
	fmt.Printf("Boolean Type: %v \n", boolZero)
	fmt.Printf("Integer Type: %v \n", intZero)
	fmt.Printf("Float Type: %v \n", floatZero)
	fmt.Printf("String Type: %q \n", strZero) // %q is Quoted String as %s and %v prints nothing as the self/value is empty.
}
