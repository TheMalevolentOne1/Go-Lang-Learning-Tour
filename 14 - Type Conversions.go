package TypeConversions

import "fmt"

func main() {
	var (
		i int = 42
		// f float64 = i // Implicit Type Conversion NOT ALLOWED
		f float64 = float64(i) // Explicit Type Conversion
		u uint    = uint(f)
	)

	fmt.Println(i, f, u)
}
