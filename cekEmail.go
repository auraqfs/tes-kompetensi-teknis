package main

import (
	"fmt"
	"strings"
)

func main() {

	var email string

	fmt.Print("Cek email : ")
	fmt.Scanf("%s", &email)

	splitEmail := strings.Split(email, "@")
	length := len([]rune(splitEmail[0]))

	if strings.Contains(email, "@") {
		if strings.Contains(splitEmail[1], ".") {
			if length <= 20 {
				if strings.Contains(splitEmail[1], ".co.id") || strings.Contains(splitEmail[1], ".id") {

					fmt.Println("Email valid")
				} else {
					fmt.Println("Domain yang diperbolehkan adalah .co.id dan .id")
				}

			} else {
				fmt.Println("Maksimal karakter sebelum @ adalah 20")
			}
		} else {
			fmt.Println("Email harus mengandung . setelah @")
		}
	} else {
		fmt.Println("Email wajib menggunakan @")
	}

}
