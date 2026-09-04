package BasicTypes

import (
	"fmt"
)

// All Exported except number
var (
	Status      bool    // Boolean
	Temperature float32 // Float32
	Balance     uint64  // Unsigned (No Negatives) Integer64
	number      byte    // Alias for uint8
)

// Printf  (Like in C) uses Format Verbs unlike python's fstring String Interpolation
func main() {
	fmt.Printf("Status Type: %T \n", Status)
	fmt.Printf("Temperature Type: %T \n", Temperature)
	fmt.Printf("Balance Type: %T \n", Balance)
	fmt.Printf("Number Type: %T \n", number)
}
