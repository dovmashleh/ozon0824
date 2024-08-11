package main

import (
	"bufio"
	"fmt"
	"os"
)

type Resource struct {
	x   int
	y   int
	typ int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var tests int
	var n, k, m int

	fmt.Fscan(in, &tests)
	for t := 0; t < tests; t++ {
		mpX := make(map[int][]*Resource, n)
		mpY := make(map[int][]*Resource, m)
		fmt.Fscan(in, &n, &m)
		fmt.Fscan(in, &k)
		var cnt, x, y int
		/*debugR := make([][]string, n)
		for i, _ := range debugR {
			debugR[i] = make([]string, m)
		}*/
		resources := make([][3]int, 0)
		for r := 0; r < k; r++ {
			fmt.Fscan(in, &cnt)
			for s := 0; s < cnt; s++ {
				fmt.Fscan(in, &x, &y)
				x--
				y--
				res := &Resource{x, y, r}
				//debugR[x][y] = debugR[x][y] + strconv.Itoa(r)
				mpX[x] = append(mpX[x], res)
				mpY[y] = append(mpY[y], res)
				resources = append(resources, [3]int{x, y, k})
			}
		}
		minS := n * m
		if n < m {
			xBeg := 0
			xEnd := -1
			hasX := false
			curHasX := make([]int, k)
			for {
				if xBeg == n {
					break
				}

				xEnd++
				if xEnd == n {
					xBeg++
					curHasX = make([]int, k)
					xEnd = xBeg - 1
					continue
				}
				for _, res := range mpX[xEnd] {
					curHasX[res.typ]++
				}

				hasX = true
				for r := 0; r < k; r++ {
					if curHasX[r] == 0 {
						hasX = false
						break
					}
				}
				if hasX {
					hasY := false
					curHasY := make([]int, k)
					yBeg := 0
					yEnd := -1
					for {
						if yEnd == m-1 && !hasY {
							break
						}
						if !hasY {
							yEnd++
							for _, res := range mpY[yEnd] {
								if res.x >= xBeg && res.x <= xEnd {
									curHasY[res.typ]++
								}
							}
						} else {
							for _, res := range mpY[yBeg] {
								if res.x >= xBeg && res.x <= xEnd {
									curHasY[res.typ]--
								}
							}
							yBeg++
						}
						hasY = true
						for r := 0; r < k; r++ {
							if curHasY[r] == 0 {
								hasY = false
								break
							}
						}
						if hasY {
							curS := (xEnd - xBeg + 1) * (yEnd - yBeg + 1)
							if curS < minS {
								minS = curS
							}
						}
					}
				}
			}
		} else {
			yBeg := 0
			yEnd := -1
			hasY := false
			curHasY := make([]int, k)
			for {
				if yBeg == m {
					break
				}

				yEnd++
				if yEnd == m {
					yBeg++
					curHasY = make([]int, k)
					yEnd = yBeg - 1
					continue
				}
				for _, res := range mpY[yEnd] {
					curHasY[res.typ]++
				}

				hasY = true
				for r := 0; r < k; r++ {
					if curHasY[r] == 0 {
						hasY = false
						break
					}
				}
				if hasY {
					hasX := false
					curHasX := make([]int, k)
					xBeg := 0
					xEnd := -1
					for {
						if xEnd == n-1 && !hasX {
							break
						}
						if !hasX {
							xEnd++
							for _, res := range mpX[xEnd] {
								if res.y >= yBeg && res.y <= yEnd {
									curHasX[res.typ]++
								}
							}
						} else {
							for _, res := range mpX[xBeg] {
								if res.y >= yBeg && res.y <= yEnd {
									curHasX[res.typ]--
								}
							}
							xBeg++
						}
						hasX = true
						for r := 0; r < k; r++ {
							if curHasX[r] == 0 {
								hasX = false
								break
							}
						}
						if hasX {
							curS := (yEnd - yBeg + 1) * (xEnd - xBeg + 1)
							if curS < minS {
								minS = curS
							}
						}
					}
				}
			}
		}
		fmt.Println("Input: n=", n, ",m=", m, ",k=", k, "deposits=", resources)
		fmt.Println("Output:", minS)
	}

}
