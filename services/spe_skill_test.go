package services

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type ISpeSkill interface {
	Narcistic(int) bool
	FindOutlier([]int) (bool, int)
	FindNeedle([]string, string) int
	BlueOceanReverse([]int, []int) []int
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

func (svc SpeSkillTest) Narcistic(number int) bool {
	stringNumber := strconv.Itoa(number)
	charArr := strings.Split(stringNumber, "")
	totalSum := 0
	times := len(charArr)
	for _, char := range charArr {
		digit, err := strconv.Atoi(char)
		if err != nil {
			return false
		}
		totalSum += powerOfNumber(digit, times)
	}

	fmt.Println(totalSum == number)
	return totalSum == number
}

func (svc SpeSkillTest) FindOutlier(numbers []int) (bool, int) {
	evens := []int{}
	odds := []int{}

	for _, number := range numbers {
		if number%2 == 0 {
			evens = append(evens, number)
		} else {
			odds = append(odds, number)
		}
	}

	valid := false
	outlier := -1

	if len(evens) == 0 || len(odds) == 0 {
		fmt.Println("false")
	} else if len(odds) > 1 && len(evens) > 1 {
		fmt.Println("false")
	} else if len(odds) == 1 {
		valid = true
		outlier = odds[0]
		fmt.Println(outlier)
	} else if len(evens) == 1 {
		valid = true
		outlier = evens[0]
		fmt.Println(outlier)
	}

	return valid, outlier
}

func (svc SpeSkillTest) FindNeedle(haystack []string, needle string) int {
	result := 0
	for index, hay := range haystack {
		if hay == needle {
			result = index
			break
		}
	}

	fmt.Println(result)
	return result
}

func (svc SpeSkillTest) BlueOceanReverse(firstList []int, secondList []int) []int {
	result := firstList
	for _, number := range secondList {
		if slices.Contains(result, number) {
			result = removeFromArray(result, number)
		}
	}
	fmt.Println(result)
	return result
}

func powerOfNumber(number int, times int) int {
	result := number
	for i := 1; i < times; i++ {
		result *= number
	}

	return result
}

func removeFromArray(items []int, item int) []int {
	newitems := []int{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
