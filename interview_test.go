package main

import (
	"fmt"
	"testing"
)

//Difficulty: easy
//Time Taken: 5m
//Time Taken: 5m
func TestTestValidity (t *testing.T) {
	result := testValidity(`1-abc-2-efg-3-hij-4-klm`)
	if !result {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(result), `true`);
	}
	result1 := testValidity(`1- -2-efg-3-hij-4-klm`)

	if !result {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(result1), `false`);
	}

	result2 := testValidity(`5-abc-2-efg-3-hij-4-klm`)
	if !result {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(result2), `false`);
	}
}

//Difficulty: easy
//Time Taken: 2m
//Time Taken: 2m
func TestWholeStory (t *testing.T) {
	result := wholeStory(`1-abc-2-efg-3-hij-4-klm`)
	if result != `abc efg hij klm` {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(result), `abc efg hij klm`);
	}
}

//Difficulty: easy
//Time Taken: 2m
//Time Taken: 2m
func TestAverageNumber (t *testing.T) {
	result := averageNumber(`1-abc-2-efg-3-hij-4-klm`)
	if result != 2 {
		t.Errorf(`TestValidity was incorrect, got: %d, want: %d.`, result, 2);
	}
}

func TestStoryStats (t *testing.T) {
	shortestArr, longestArr, averageLen, averageStrings := storyStats(`1-abc-2-efgh-3-ijklm-4-nopqrst`)

	return1CheckPoint := []string{`abc`}
	if !testEq(shortestArr, return1CheckPoint) {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(shortestArr), fmt.Sprint(return1CheckPoint))
	}
		return2CheckPoint := []string{`nopqrst`}
	if !testEq(longestArr, return2CheckPoint) {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(longestArr), fmt.Sprint(return2CheckPoint))
	}

	if averageLen != 4 {
		t.Errorf(`TestValidity was incorrect, got: %d, want: %d.`, averageLen, 3);
	}

	return4CheckPoint := []string{`ijklm`, `efgh`}
	if !testEq(averageStrings, return4CheckPoint) {
		t.Errorf(`TestValidity was incorrect, got: %s, want: %s.`, fmt.Sprint(averageStrings), fmt.Sprint(return4CheckPoint))
	}
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
