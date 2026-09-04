package Constants

import (
	"fmt"
	"math"
)

const pi float64 = 3.141592653589793

func main() {
	// Calculate Circle Area
	// πr2 (Radius is 12)
	fmt.Println(pi * math.Pow(12, 2))
}
