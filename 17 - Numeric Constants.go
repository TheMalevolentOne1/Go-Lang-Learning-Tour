package NumericConstants

import "fmt"

const (
	LargeInt = 1 << 100
	SmallInt = LargeInt >> 99
)

func convertInt(x int) int           { return x*10 + 1 }
func convertFloat(x float64) float64 { return x * 0.1 } // no idea why the times and addition was on the tour.

func main() {
	// fmt.Println(convertInt(LargeInt)) Due to size of value causes the unsigned constant when converted to int64 to overflow, therefore compiler refuses.
	fmt.Println(convertInt(SmallInt))
	fmt.Println(convertFloat(LargeInt))
	fmt.Println(convertFloat(SmallInt))
}
