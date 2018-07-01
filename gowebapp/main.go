package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"app/simpleweb"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	//runtime.GOMAXPROCS(1)
	// fmt.Println(mymath.Sqrt(77.0))
	// go say("hello")
	// say("world")

	// a := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(a[:len(a)/2], c)
	// go sum(a[len(a)/2:], c)

	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// cBuf := make(chan int, 2)
	// cBuf <- 1
	// cBuf <- 2
	// fmt.Println(<-cBuf)
	// fmt.Println(<-cBuf)
	// fmt.Println(cap(cBuf))
	// fmt.Println("Starting fibonacci sequence")
	// cFib, quit := make(chan int), make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-cFib)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci(cFib, quit)
	http.HandleFunc("/", simpleweb.SayhelloName)
	http.HandleFunc("/zamir", simpleweb.SayhelloNameToZamir)
	http.HandleFunc("/login", simpleweb.Login)
	http.HandleFunc("/upload", simpleweb.Upload)
	err := http.ListenAndServe(":"+PORT, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Printf("Listening at http://localhost:%s ...", PORT)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
