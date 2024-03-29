package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func proccess(teks string, characters string, angka string) string{ // logic pass
	data1 := strings.Split(teks, "") //  [a b c]

	output := []string{"", "", "", "", "", "", ""}

	for i, v := range data1 {
		random := rand.Float32()  // 0.3412831238
		bulat := math.Round(float64(random)) // 1 atau 0
		// fmt.Println(bulat)
		if bulat == 1 {
			temp := strings.ToUpper(v)
			output[i] = temp
		} else {
			temp := strings.ToLower(v)
			output[i] = temp
		}
	}
	// AbC
	
	gabung := strings.Join(output, "")
	return (gabung + characters + angka) // output
}

func GenPass(pass string, level string) string{
	if len(pass) <= 5 {
		return "Password Not Strong!"
	} else if len(pass) > 7 {
		return "Password Over load...\nPassword must be 6 - 7 characters !"
	}
	random := rand.Perm(9) // mencari random angka [1 2 3 4 5 9]
	angka := 3
	if level == "medium" {
		angka = 4
	} else if level == "strong" {
		angka = 5
	}
	hasilAngka := random[0:angka] // [3 4 2]
	outputAngka := strings.Trim(strings.Replace(fmt.Sprint(hasilAngka), " ", "", -1), "[]") // 342

	specialCharacters := "@:;#&-?/%+*" // mencari random character
	special := strings.Split(specialCharacters, "") // [@ : ; # & - ? / % + *] 11
	panjang := len(special)
	getspesial := rand.Intn(panjang) // 2
	char := "" // ;

	for i, v := range special{
		if i == getspesial {
			char = v
		}
	}
	
	return proccess(pass, char, outputAngka)
}