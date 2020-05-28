// https://blog.golang.org/strings
// https://gobyexample.com/http-servers

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	resp, err := http.Get("http://terokarvinen.com/2020/go-programming-course-2020-w22")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	var t []int
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		if (rand.Intn(2-1+1) + 1) == 2 {
			fmt.Println("rivi :", i+1, " ", scanner.Text())

			t = append(t, len(scanner.Text()))
		}
	}
	fmt.Println("Rivien merkkipituudet:", t)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
