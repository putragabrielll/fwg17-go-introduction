package main

import (
	"fmt"
)


func MainData(n int) string{
	nilai := []int{1, 7, 3, 4, 8, 9}

	temp := 0
	
	for _, data1 := range nilai {
		// fmt.Println(n)
		for _, data2 := range nilai {
			temp = data1 + data2

			if temp == n {
				hasil := fmt.Sprintf("%v dan %v", data1, data2)
				return hasil
			}
		}
	}

	return fmt.Sprintln("Not Found!")
}