package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func proccess(teks string, characters string){ // logic pass
	data1 := strings.Split(teks, "") //  [a b c]

	output := []string{"", "", "", ""}

	for i, v := range data1 {
		random := rand.Float32()
		bulat := math.Round(float64(random))
		// fmt.Println(bulat)
		if bulat == 1 {
			temp := strings.ToUpper(v)
			output[i] = temp
		} else {
			temp := strings.ToLower(v)
			output[i] = temp
		}
	}
	
	gabung := strings.Join(output, "")
	fmt.Println(gabung + characters) // output
}

func GenPass(pass string, level string){
	specialCharacters := "@:;#&-?/%+*"
	special := strings.Split(specialCharacters, "") // [@ : ; # & - ? / % + *]
	panjang := len(special)
	getspesial := rand.Intn(panjang)
	char := ""

	for i, v := range special{
		if i == getspesial {
			char = v
		}
	}

	proccess(pass, char)
}