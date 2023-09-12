package main

import "github.com/chandra-phang/spe-skill-test/services"

func main() {
	services.InitSpeSkillTest()
	speSkillTest := services.GetSpeSkillTest()

	speSkillTest.Narcistic(153)
	speSkillTest.Narcistic(111)
	speSkillTest.FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})
	speSkillTest.FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21})
	speSkillTest.FindOutlier([]int{11, 13, 15, 19, 9, 13, -21})
	speSkillTest.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue")
	speSkillTest.BlueOceanReverse([]int{1, 2, 3, 4, 6, 10}, []int{1})
	speSkillTest.BlueOceanReverse([]int{1, 5, 5, 5, 5, 3}, []int{5})
}
