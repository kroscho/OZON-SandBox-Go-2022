package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func inputInts(in *bufio.Reader, text string) (int, int) {
	var x, y int
	fmt.Print(text)
	fmt.Fscan(in, &x, &y)

	return x, y
}

func inputInt(in *bufio.Reader, text string) int {
	var x int
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

type Proccess struct {
	isBusy  bool
	Energy  int
	EndTime int
}

func NewProccess(energy int) *Proccess {
	return &Proccess{
		isBusy:  false,
		Energy:  energy,
		EndTime: 1,
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	countProccess, countTasks := inputInts(in, "Введите кол-во процессов и задач: ")
	proccesses := make([]*Proccess, countProccess)

	for i := 0; i < countProccess; i++ {
		en := inputInt(in, "Введите энергопотребление: ")
		proccesses[i] = NewProccess(en)
	}

	sort.Slice(proccesses, func(i, j int) (less bool) {
		return proccesses[i].Energy < proccesses[j].Energy
	})

	t := 1
	sum := 0
	for i := 0; i < countTasks; i++ {
		timeStart, timeDuration := inputInts(in, "\nВведите время старта и продолжительность: ")

		for j := t; j <= timeStart; j++ {
			for id, p := range proccesses {
				if p.isBusy && p.EndTime == j {
					fmt.Printf("\nОсвобождается %d процессор.", id+1)
					p.isBusy = false
				}
			}
		}

		allBusy := true
		for id, p := range proccesses {
			if !p.isBusy {
				p.EndTime = timeStart + timeDuration
				p.isBusy = true
				sum += p.Energy * timeDuration
				fmt.Printf("\nПришла %d задача. Занимает %d процессор.", i+1, id+1)
				allBusy = false
				break
			}
		}
		if allBusy {
			fmt.Printf("\nПришла %d задача. Все процессоры заняты, задача отбрасывается.", i+1)
		}
		t += 1
	}
	fmt.Println("\n\nSum: ", sum)
}
