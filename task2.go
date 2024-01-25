package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func proccess(teks string, characters string, angka string) string{ // logic pass
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
	return (gabung + characters + angka) // output
}

func GenPass(pass string, level string) string{
	random := rand.Perm(9)
	angka := 3
	if level == "medium" {
		angka = 4
	} else if level == "strong" {
		angka = 5
	}
	hasilAngka := random[0:angka]
	outputAngka := strings.Trim(strings.Replace(fmt.Sprint(hasilAngka), " ", "", -1), "[]")

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
	
	return proccess(pass, char, outputAngka)
}