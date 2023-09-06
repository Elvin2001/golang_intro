package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i == 10 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------------------------------------------------")

	b := 10
	switch b {
	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")

	default:
		fmt.Println("Not 1 and 2")
	}

	fmt.Println("-------------------------------------------------")

	defer fmt.Println("Hello!")
	fmt.Println("Goodbye!")

}
