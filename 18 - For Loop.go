package ForLoop

import "fmt"

// For Loops
func forStatementOne() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forStatementTwo() {
	i := 0
	for i < 10 {
		i = i + 2
	}
	fmt.Println(i)
}

// While Loop
func forStatementThree() {
	// For Loop
	i := 0
	for ; i < 10; i += 2 {
		fmt.Println(i)
	}

	// While Loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j = j + 2
	}
}

func forStatementFour() {
	for {
	}
}
