package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

/*
func check(e error) {
	if e != nil {
		panic(e)
	}
}
*/
func main() {
	var text string
	flag.StringVar(&text, "text", "default", "Use --text to enter text to be written in a file")
	flag.Parse()

	t := []byte(text)
	err := ioutil.WriteFile("./generoitu.txt", t, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	//check(err)
}
