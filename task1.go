package main

import "fmt"

func printSegitiga(angka int) {
	// var k int
	for initial := 1; angka >= initial; angka-- {
		// k = angka
		// angka = 5
		for i := 1; i <= angka; i++ {
			fmt.Print(angka)
		}

		for j := 1; j <= angka - initial; j++ {
			fmt.Print(" ")
		}

		// for {
		// 	fmt.Print("* ")
		// 	k--
		// 	if k == 2 * initial - 1 {
		// 		break
		// 	}
		// }
		fmt.Println("")
	}
}