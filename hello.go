package main

import "fmt"

func main() {
	var a = "hello world"
	var b = "hello"
	var c = "world"
	var bil int

	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&bil)

	if (bil%3 == 0) && (bil%5 == 0) {
		fmt.Println(a)
	} else if bil%3 == 0 {
		fmt.Println(b)
	} else if bil%5 == 0 {
		fmt.Println(c)
	} else {
		fmt.Println("Masukan bilangan kembali")
	}

}
