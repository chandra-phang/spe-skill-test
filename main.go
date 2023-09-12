package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ISpeSkill interface {
	NarcisticNumber(int) bool
}

type SpeSkillTest struct {
}

var speSingleton ISpeSkill

func InitSpeSkillTest() {
	speSingleton = SpeSkillTest{}
}

func GetSpeSkillTest() ISpeSkill {
	return speSingleton
}

func (svc SpeSkillTest) NarcisticNumber(number int) bool {
	stringNumber := strconv.Itoa(number)
	charArr := strings.Split(stringNumber, "")
	totalSum := 0
	times := len(charArr)
	for _, char := range charArr {
		digit, err := strconv.Atoi(char)
		if err != nil {
			return false
		}
		fmt.Printf("powerOfNumber(%d, %d): %d", digit, times, powerOfNumber(digit, times))
		totalSum += powerOfNumber(digit, times)
	}
	return totalSum == number
}

func powerOfNumber(number int, times int) int {
	result := number
	for i := 1; i < times; i++ {
		result *= number
	}

	return result
}

func main() {
	InitSpeSkillTest()
	speSkillTest := GetSpeSkillTest()
	speSkillTest.NarcisticNumber(153)
}
