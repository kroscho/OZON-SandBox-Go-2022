package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	priceMap := make(map[int]int, 0)

	for i := 0; i < testsCount; i++ {
		n := inputInt(in, "Введите кол-во товаров: ")

		for j := 0; j < n; j++ {
			price := inputInt(in, "Введите цену: ")
			priceMap[price] += 1
		}

		sum := 0
		for price, count := range priceMap {
			countThree := count / 3
			if count != 0 {
				count -= countThree
			}
			sum += price * count
		}
		fmt.Print("Сумма с учетом акции: ")
		fmt.Fprintln(out, sum)
	}
}
