package example

import (
	"fmt"
	"math"
	"time"
)

func Forloop(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum
}

func Forsim(num int) int {
	if num < 1 {
		fmt.Printf("value of %d is not proper\n", num)
		return num
	}
	sum := 1
	for sum < num {
		sum += 2
	}
	fmt.Println(sum)
	return sum
}

func Forw(num int) int {
	sum := 1
	for sum < num {
		sum += sum
	}
	fmt.Println(sum)
	return sum
}

func Calpow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g", v, lim)
	}
	return lim
}

func Gettoday() {
	today := time.Now().Weekday()
	fmt.Println("Today is ", today)
	fmt.Printf("Saturday is ")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}

func Defergame() {
	fmt.Println("defer game starting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("defer game end")

}
