package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	if x := 4; x%2 == 0 {
		fmt.Println(x,"is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is Negative")
	} else if num<10 {
		fmt.Println(num,"has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
