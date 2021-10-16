package main

import (
	"fmt"
	"math"
)

const n = 20

func main() {
	fmt.Println("hello world")
	fmt.Println("hello he is " + "good boy ")
	fmt.Println(2 + 3)
	fmt.Println((true && false))
	var a = "initial"
	fmt.Println(a)
	var b, c int = 20, 21
	fmt.Println(b, c)
	var e int
	fmt.Println(e)
	m := "alan"
	fmt.Println(m)
	f := "apple"
	fmt.Println(f)
	const h = "jack"
	fmt.Println(h)
	fmt.Println(math.Sin(n))
	i := 0
	for i <= 20 {
		fmt.Println(i)
		i = i + 1

	}
	fmt.Println(i % 21)

	if i%21 == 0 {
		fmt.Println("no reached 20")
	} else if i < 20 {
		fmt.Println("7 is odd")

	} else {
		fmt.Println("helo last one")
	}
	sam := 2
	switch sam {
	case 1:
		fmt.Println("one is here")
	case 2:
		fmt.Println(("commcected"))
	case 3:
		fmt.Println("hey i am here")
	default:
		fmt.Println("everything ends here")
	}

}
