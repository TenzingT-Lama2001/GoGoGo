/**

Large programs are often made up of many smaller sub-programs.
Eg. a web server handles requests made from web browsers and serves up HTML web pages
in response each request is handled like a small program.

It would be ideal for prgrams like theseto beable to run their smaller components at the
same time(in the case of the web server to hanlde multiple requests). Making progress on more
than one task simultaneously is known as concurrency.

GO has rich support for concurrency using go routines and channels.

**/

/** Go routines

A goroutine is a function that is capable of running concurrentyl with other functions.
TO reate a goroutine we use the keyword "go" followed by a function invocation:

**/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 		amt := time.Duration(rand.Intn(250)) // random time between 0 and 250 milliseconds
// 		time.Sleep(time.Millisecond * amt)
// 	}
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go f(i)
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }

// func main() {
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

var wg sync.WaitGroup //pointer

func getStatusCode(endpoint string) {
	defer wg.Done() //decrement the wait group counter
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%d staus code for %s\n", res.StatusCode, endpoint)
	}

}

func main() {

	websiteList := []string{
		"https://lco.dev",
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com"}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) //add 1 to the wait group
	}

	wg.Wait() //wait for all the go routines to finish

}
