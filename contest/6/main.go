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

	var tests int
	var n int

	fmt.Fscan(in, &tests)
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &n)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &(nums[i]))
		}
		diffsWithPrev := make(map[int]int)
		pars := 0
		for i := 1; i < n-1; i++ {
			diffPost := nums[i] - nums[i+1]
			if cnt, ok := diffsWithPrev[diffPost]; ok {
				pars += cnt
			}
			diffPrev := nums[i-1] - nums[i]
			diffsWithPrev[diffPrev]++
		}
		fmt.Fprintln(out, pars)
	}
}
