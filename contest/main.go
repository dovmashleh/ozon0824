package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// -------------------------------
	var countPacks int
	fmt.Fscan(in, &countPacks)

	for i := 0; i < countPacks; i++ {
		var countNums int
		fmt.Fscan(in, &countNums)

		// from 1 to n
		nums := make([]int, 0, countNums)
		for j := 0; j < countNums; j++ {
			var num int
			fmt.Fscan(in, &num)
			nums = append(nums, num)
		}

		sums := make(map[int][]struct{ a, b int }, countNums)
		for a := 0; a < countNums-1; a++ {
			for b := a + 1; b < countNums; b++ {
				sum := nums[a] + nums[b]
				sums[sum] = append(sums[sum], struct{ a, b int }{a, b})
			}
		}

		result := make(map[struct{ a, b int }]struct{})
		for _, vars := range sums {
			if len(vars) < 2 {
				continue
			}

			for a := 0; a < len(vars)-1; a++ {
				for b := a + 1; b < len(vars); b++ {
					if vars[a].a+1 == vars[b].a && vars[a].b-1 == vars[b].b {
						result[struct{ a, b int }{
							a: vars[b].a, b: vars[b].b,
						}] = struct{}{}
					}
					if vars[b].a+1 == vars[a].a && vars[b].b-1 == vars[a].b {
						result[struct{ a, b int }{
							a: vars[b].a, b: vars[b].b,
						}] = struct{}{}
					}
				}

			}
		}

		fmt.Fprintln(out, len(result))
	}
}
