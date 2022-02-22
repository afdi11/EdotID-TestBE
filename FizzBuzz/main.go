package main

import (
	"EdotID-TestBE/FizzBuzz/generate"
	"fmt"
)

func main() {
	var angka int
	var answer []string
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)
	answer = generate.Answer(angka)
	// fmt.Println(answer)
	generate.Cetak(answer)
}
