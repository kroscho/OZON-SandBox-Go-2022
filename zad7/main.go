package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	ID        int
	Color     string
	Neighbors []int
	Visited   bool
}

func NewVertex(id int, color string, m int, countVertex int, isLeft bool, isRight bool, row int) Vertex {
	leftShift := 0
	rightShift := 0

	neighbors := []int{}
	// поиск соседей
	if m%2 == 0 {
		leftShift = m / 2
		rightShift = m/2 + 1

		if isLeft && isRight {
			topNeigbor1 := id - 2
			topNeigbor2 := id - 1
			bottomNeighbor1 := id + 1
			bottomNeighbor2 := id + 2
			if topNeigbor1 >= 0 {
				neighbors = append(neighbors, topNeigbor1)
				neighbors = append(neighbors, topNeigbor2)
			} else if topNeigbor2 >= 0 {
				neighbors = append(neighbors, topNeigbor2)
			}

			if bottomNeighbor1 < countVertex {
				neighbors = append(neighbors, bottomNeighbor1)
				neighbors = append(neighbors, bottomNeighbor2)
			} else if bottomNeighbor2 < countVertex {
				neighbors = append(neighbors, bottomNeighbor2)
			}
		} else {
			if isLeft && row%2 != 0 {
				topNeigbor := id - leftShift
				bottomNeighbor := id + (rightShift - 1)
				if topNeigbor >= 0 {
					neighbors = append(neighbors, topNeigbor)
				}

				neighbors = append(neighbors, id+1)

				if bottomNeighbor < countVertex {
					neighbors = append(neighbors, bottomNeighbor)
				}
			} else if isLeft && row%2 == 0 {
				topNeigbor1 := id - leftShift
				topNeigbor2 := id - (leftShift - 1)
				bottomNeighbor1 := id + (rightShift - 1)
				bottomNeighbor2 := id + rightShift
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id+1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			}
			if isRight && row%2 != 0 {
				topNeigbor1 := id - (leftShift + 1)
				topNeigbor2 := id - leftShift
				bottomNeighbor1 := id + (rightShift - 1)
				bottomNeighbor2 := id + (rightShift - 2)
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id-1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			} else if isRight && row%2 == 0 {
				topNeigbor1 := id - leftShift
				bottomNeighbor1 := id + (rightShift - 1)
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
				}

				neighbors = append(neighbors, id-1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
				}
			}
			if !isLeft && !isRight {
				topNeigbor1 := 0
				topNeigbor2 := 0
				bottomNeighbor1 := 0
				bottomNeighbor2 := 0
				if row%2 != 0 {
					topNeigbor1 = id - leftShift
					topNeigbor2 = id - (leftShift - 1)
					bottomNeighbor1 = id + (rightShift - 1)
					bottomNeighbor2 = id + (rightShift - 2)
				} else {
					topNeigbor1 = id - leftShift
					topNeigbor2 = id - (leftShift - 1)
					bottomNeighbor1 = id + (rightShift - 1)
					bottomNeighbor2 = id + rightShift
				}
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id-1)
				neighbors = append(neighbors, id+1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			}
		}
	} else {
		leftShift = m/2 + 1
		rightShift = m/2 + 1

		if isLeft && isRight {
			topNeigbor1 := id - 2
			topNeigbor2 := id - 1
			bottomNeighbor1 := id + 1
			bottomNeighbor2 := id + 2
			if topNeigbor1 >= 0 {
				neighbors = append(neighbors, topNeigbor1)
				neighbors = append(neighbors, topNeigbor2)
			} else if topNeigbor2 >= 0 {
				neighbors = append(neighbors, topNeigbor2)
			}

			if bottomNeighbor1 < countVertex {
				neighbors = append(neighbors, bottomNeighbor1)
				neighbors = append(neighbors, bottomNeighbor2)
			} else if bottomNeighbor2 < countVertex {
				neighbors = append(neighbors, bottomNeighbor2)
			}
		} else {
			if isLeft && row%2 != 0 {
				topNeigbor := id - (leftShift - 1)
				bottomNeighbor := id + rightShift
				if topNeigbor >= 0 {
					neighbors = append(neighbors, topNeigbor)
				}

				neighbors = append(neighbors, id+1)

				if bottomNeighbor < countVertex {
					neighbors = append(neighbors, bottomNeighbor)
				}
			} else if isLeft && row%2 == 0 {
				topNeigbor1 := id - leftShift
				topNeigbor2 := id - (leftShift - 1)
				bottomNeighbor1 := id + rightShift
				bottomNeighbor2 := id + (rightShift - 1)
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id+1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			}
			if isRight && row%2 != 0 {
				topNeigbor := id - leftShift
				bottomNeighbor := id + (rightShift - 1)
				if topNeigbor >= 0 {
					neighbors = append(neighbors, topNeigbor)
				}

				neighbors = append(neighbors, id-1)

				if bottomNeighbor < countVertex {
					neighbors = append(neighbors, bottomNeighbor)
				}
			} else if isRight && row%2 == 0 {
				topNeigbor1 := id - leftShift
				topNeigbor2 := id - (leftShift - 1)
				bottomNeighbor1 := id + rightShift
				bottomNeighbor2 := id + (rightShift - 1)
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id-1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			}
			if !isLeft && !isRight {
				topNeigbor1 := id - leftShift
				topNeigbor2 := id - (leftShift - 1)
				bottomNeighbor1 := id + rightShift
				bottomNeighbor2 := id + (rightShift - 1)
				if topNeigbor1 >= 0 {
					neighbors = append(neighbors, topNeigbor1)
					neighbors = append(neighbors, topNeigbor2)
				} else if topNeigbor2 >= 0 {
					neighbors = append(neighbors, topNeigbor2)
				}

				neighbors = append(neighbors, id-1)
				neighbors = append(neighbors, id+1)

				if bottomNeighbor1 < countVertex {
					neighbors = append(neighbors, bottomNeighbor1)
					neighbors = append(neighbors, bottomNeighbor2)
				} else if bottomNeighbor2 < countVertex {
					neighbors = append(neighbors, bottomNeighbor2)
				}
			}
		}
	}

	return Vertex{
		ID:        id,
		Color:     color,
		Visited:   false,
		Neighbors: neighbors,
	}
}

type T struct{}

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

func inputStr(in *bufio.Reader, text string) string {
	var x string
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

func findNeighboar(curVertex *Vertex, vertexes []*Vertex) {
	colorExist := false

	for _, neighborID := range curVertex.Neighbors {
		if vertexes[neighborID].Visited {
			continue
		}
		if vertexes[neighborID].Color == curVertex.Color {
			vertexes[neighborID].Visited = true
			colorExist = true
			findNeighboar(vertexes[neighborID], vertexes)
		}
	}
	if !colorExist {
		return
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	countTests := inputInt(in, "Введите кол-во тестов: ")

	for i := 0; i < countTests; i++ {
		n, m := inputInts(in, "Введите кол-во строк и символов в строке: ")
		countVertex := n * m / 2
		if (n*m)%2 != 0 {
			countVertex += 1
		}
		vertexMap := make(map[int]Vertex, countVertex)
		vertexes := make([]*Vertex, countVertex)

		idVertex := 0

		for j := 0; j < n; j++ {
			row := inputStr(in, "Введите строку: ")
			row = strings.TrimLeft(row, ".")
			row = strings.TrimRight(row, ".")
			colors := strings.Split(row, ".")

			for k := 0; k < len(colors); k++ {
				isLeft := false
				isRight := false
				if k == 0 {
					isLeft = true
				}
				if k == len(colors)-1 {
					isRight = true
				}
				v := NewVertex(idVertex, colors[k], m, countVertex, isLeft, isRight, j+1)
				vertexes[idVertex] = &v
				vertexMap[v.ID] = v
				idVertex += 1
			}
		}

		colorsMap := make(map[string]T, 0)

		for _, v := range vertexes {
			if v.Visited {
				continue
			}
			_, ok := colorsMap[v.Color]
			if ok {
				fmt.Fprintln(out, "NO")
				return
			}
			colorsMap[v.Color] = T{}
			v.Visited = true
			findNeighboar(v, vertexes)
		}
		fmt.Fprintln(out, "YES")
	}
}
