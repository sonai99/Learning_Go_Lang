package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// fmt.Print(time.Now(), time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("ITs the Weekend")
	default:
		fmt.Println("ITs a Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("ITs after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("Dont know the type of ", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
