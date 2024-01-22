package main

import "fmt"

func printSegitiga(angka int) { // 5
	var p int = 1
	for initial := 1; initial <= angka; initial++ {
		if initial != 1 {
			// cara 1
			// for i := 1; i <= angka-(angka - initial + 1); i++ {
			// 	fmt.Print("  ")
			// }

			// cara 2
			for i := 1; i <= p; i++ {
				fmt.Print("  ")
			}
			p++
		}

		// var h int = angka * 2 - 1
		// if h == angka {
		// 	fmt.Print(("  "))
		// }

		for i := 1; i <= angka - initial + 1; i++ {
			fmt.Print("* ")	
		}
		for j := 1; j <= angka - initial; j++ {
			fmt.Print("* ")
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