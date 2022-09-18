package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct{}

func inputInt(in *bufio.Reader, text string) int {
	var x int
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	testsCount := inputInt(in, "Введите количество тестов: ")

	tasksMap := make(map[int]Task)

	for i := 0; i < testsCount; i++ {
		daysCount := inputInt(in, "Введите количество дней: ")
		lastTask := inputInt(in, "Введите задание: ")

		tasksMap[lastTask] = Task{}

		isYes := true
		for j := 0; j < daysCount-1; j++ {
			task := inputInt(in, "Введите задание: ")

			if task == lastTask {
				continue
			}
			_, ok := tasksMap[task]
			if ok {
				fmt.Fprintln(out, "NO")
				isYes = false
				break
			} else {
				tasksMap[task] = Task{}
				lastTask = task
			}
		}
		if isYes {
			fmt.Fprintln(out, "YES")
		}
	}
}
