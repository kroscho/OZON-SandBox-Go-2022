package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Developer struct {
	Index int
	Skill int
}

func inputInt(in *bufio.Reader, text string) int {
	var x int
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

func removeByIndex(array []Developer, index int) []Developer {
	return append(array[:index], array[index+1:]...)
}

func getDefferenceSkills(x1 int, x2 int) float64 {
	return math.Abs(float64(x1) - float64(x2))
}

func createListDevelopers(in *bufio.Reader) []Developer {
	n := inputInt(in, "Введите кол-во разработчиков: ")

	developers := make([]Developer, 0)
	for j := 0; j < n; j++ {
		skill := inputInt(in, "Введите мастерство: ")

		dev := Developer{
			Index: j + 1,
			Skill: skill,
		}
		developers = append(developers, dev)
	}

	return developers
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	testsCount := inputInt(in, "Введите количество тестов: ")

	for i := 0; i < testsCount; i++ {
		developers := createListDevelopers(in)
		for len(developers) != 0 {
			if len(developers) == 2 {
				fmt.Fprintln(out, developers[0].Index, " ", developers[1].Index)
				break
			}
			partnerIndex := 1
			min := getDefferenceSkills(developers[0].Skill, developers[1].Skill)
			for j := 2; j < len(developers); j++ {
				curDiff := getDefferenceSkills(developers[0].Skill, developers[j].Skill)
				if curDiff < min {
					partnerIndex = j
					min = curDiff
				}
			}
			fmt.Fprintln(out, developers[0].Index, " ", developers[partnerIndex].Index)
			developers = removeByIndex(developers, 0)
			developers = removeByIndex(developers, partnerIndex-1)
		}
		if i != testsCount-1 {
			fmt.Fprintln(out, "")
		}
	}
}
