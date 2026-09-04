package MultiReturns

import "fmt"

func swap(s1, s2 string) (string, string) {
	s1t := s2 // temp s1 to store s2 before overwrite
	s2 = s1
	s1 = s1t
	return s1, s2
}

func main() {
	fmt.Println(swap("Hello", "World")) // Println can only handle one return value as it's argument it cannot handle multiple return with a string.
}
