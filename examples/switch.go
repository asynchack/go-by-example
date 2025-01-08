package main

import (
	"fmt"
	"time"
)

func main()  {
	
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: 
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	// switch without an expression is an alternate way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	case t.Hour() > 12:
		fmt.Println("after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("other type")
		}
	}

	whatAmI(100)
	whatAmI(true)
	whatAmI("good")
}