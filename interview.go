package main

import (
	"math"
	"strconv"
	"strings"
)

//Difficulty: easy
//Estimate Effort: 10m
//Time Taken: 15m
func testValidity(myString string) bool {
	strArray := strings.Split(myString, `-`)
	if len(strArray) < 2 || (len(strArray) % 2) != 0 {
		return false
	}
	startNumber := 0
	for i := 0; i < len(strArray) -1; i = i + 2 {
		 if crtNumber, err := strconv.Atoi(strArray[i]); err == nil {
		 	if startNumber >= crtNumber {
				return false
			} else {
				startNumber = crtNumber
			}
		 } else {
		 	return false
		 }
		if !isLetter(strArray[i + 1]) {
			return  false
		}
	}
	return true
}

//Difficulty: easy
//Estimate Effort: 2m
//Time Taken: 2m
//supposed the string has validated, don't check it again
//not clear about the precision, return the int instead
func averageNumber(myString string) int {
	strArray := strings.Split(myString, `-`)
	var total = 0
	for i := 0; i < len(strArray) -1; i = i + 2 {
		crtNumber, _ := strconv.Atoi(strArray[i])
		total = total + crtNumber
	}
	return total/(len(strArray)/2)
}

//Difficulty: easy
//Estimate Effort: 2m
//Time Taken: 2m
//supposed the string has validated, don't check it again
func wholeStory(myString string) string {
	strArray := strings.Split(myString, `-`)
	var story string
	for i := 0; i < len(strArray) -1; i = i + 2 {
		if i != 0 {
			story = story + ` `
		}
		story = story + strArray[i+1]
	}
	return story
}

//Difficulty: easy
//Estimate Effort: 8m
//Time Taken: 25m
//supposed the string has validated, don't check it again
func storyStats (myString string) (shortestArray []string, longestArray []string, averageLen int, averageLenString []string){
	strArray := strings.Split(myString, `-`)
	shortestLen := len(strArray[1])
	longestLen := len(strArray[1])
	totalLength := 0
	lengthIndex := make(map[int][]int, 0)
	for i := 0; i < len(strArray) -1; i = i + 2 {
		currentIndex := i+1
		ctrStr := strArray[currentIndex]

		totalLength = totalLength + len(ctrStr)
		currentLen := len(ctrStr)

		if  currentLen < shortestLen {
			shortestArray = []string{ctrStr}
			shortestLen = currentLen
		} else if currentLen == shortestLen {
			shortestArray = append(shortestArray, ctrStr)
		}

		if currentLen > longestLen {
			longestArray = []string{ctrStr}
			longestLen = currentLen
		} else if currentLen == longestLen {
			longestArray = append(longestArray, ctrStr)
		}

		if lengthIndex[currentLen] == nil || len(lengthIndex[currentLen]) == 0{
			lengthIndex[currentLen] = []int{currentIndex}
		} else {
			lengthIndex[currentLen] = append(lengthIndex[currentLen], currentIndex)
		}
	}
	averageLen = totalLength/(len(strArray)/2)

	averageUp := math.Ceil(float64(totalLength)/float64(len(strArray)/2))
	averageDown := math.Floor(float64(totalLength)/float64(len(strArray)/2))
	//add round up
	if strIndexArray, exist := lengthIndex[int(averageUp)]; exist {
		for _, averageStr := range strIndexArray {
			averageLenString = append(averageLenString, strArray[averageStr])
		}
	}
	//add round down
	if strIndexArray, exist := lengthIndex[int(averageDown)]; exist {
		for _, averageStr := range strIndexArray {
			averageLenString = append(averageLenString, strArray[averageStr])
		}
	}
	return shortestArray, longestArray, averageLen,averageLenString
}


func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
