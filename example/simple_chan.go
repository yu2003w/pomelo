package example

import "fmt"

//glog.Logger()
func sum(s []int, c chan int, str string) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("%s, sum is %d\n", str, sum)
	c <- sum
}

func Gchan() {
	data := []int{1, 2, 3, 7, -8, -2}
	c := make(chan int)
	go sum(data[:len(data)/2], c, "first thread")
	go sum(data[len(data)/2:], c, "second thread")
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	// buffered channel
	bc := make(chan string, 5)
	bc <- "hello world"
	bc <- "buffered channel for go"
	// a little interesting here, with one "<-bc" couldn't receive two sends from buffered channel
	fmt.Println(<-bc)
	fmt.Println(<-bc)
	close(bc)
	v, ok := <-bc
	if !ok {
		fmt.Println("channel closed")
	} else {
		fmt.Println(v)
	}
}

func sel(in, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case in <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit received")
			return
		}
	}
}

func Gsel() {
	in := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-in)
		}
		quit <- 0
	}()
	sel(in, quit)
}
