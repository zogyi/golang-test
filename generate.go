package main

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

//Difficulty: easy
//Estimate Effort: 25m
//Time Taken: 55m
func generate(goodOrNot bool) string {
	letters := []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	notAllowed := []rune(`123456780!@#$%^&*()_+=-';`)
	length := (rand.Intn(math.MaxUint8) + 1) * 2
	stringArray := make([]string, 0)
	numberArray := make([]int, 0)
	for i := 0; i < length/2; i++ {
		crtStr := ``
		strLen := rand.Intn(math.MaxUint8) + 1
		for k := 0; k < strLen; k++ {
			if goodOrNot {
				crtStr = crtStr + string(letters[rand.Intn(51)])
			} else {
				choseArray := rand.Intn(2);
				if choseArray == 0 {
					crtStr = crtStr + string(letters[rand.Intn(52)])
				} else {
					crtStr = crtStr + string(notAllowed[rand.Intn(24)])
				}
			}

		}
		stringArray = append(stringArray, crtStr)
		numberArray = append(numberArray, rand.Intn(math.MaxUint32-1) +1)
	}
	wholeArray := make([]string, 0, length)

	if goodOrNot {
		sort.Ints(numberArray)
	}

	for i := 0; i < length /2; i++ {
		wholeArray = append(wholeArray, strconv.Itoa(numberArray[i]), stringArray[i])
	}
	return strings.Join(wholeArray, `-`)
}