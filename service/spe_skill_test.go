package service

import (
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

func (svc SpeSkillTest) NarcisticNumber(number int) bool {
	stringNumber := strconv.Itoa(number)
	charArr := strings.Split(stringNumber, "")
	totalSum := 0
	times := len(charArr)
	for _, char := range charArr {
		digit := strconv.ParseInt(char, 10, 64)
		totalSum += powerOfNumber(digit, times)
	}

	return totalSum == number
}

func powerOfNumber(number int, times int) int {
	result := number
	for i := 1; i <= times; i++ {
		result *= number
	}

	return number
}
