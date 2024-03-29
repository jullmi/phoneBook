package internal

import (
	"math/rand"
)

var MAX int = 90
var MIN int = 65



func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func generateString(len int64) string {

	temp := ""
	var count int64 = 1

	for {
		myRand := random(MIN, MAX)
		newChar := string(byte(myRand))
		temp = temp + newChar

		if count == len {
			break
		}
		count++

	}
	return temp

}



