package main

import "fmt"

func sanoMoi(i int, g string) {
	fmt.Printf("morjesta\n%v %v\n", i, g)
}

func main() {
	fmt.Println("moi")
	sanoMoi(4, "Momo")
}
