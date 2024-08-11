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
	var n, k, m int

	fmt.Fscan(in, &tests)
	for t := 0; t < tests; t++ {
		fmt.Fscan(in, &n, &k)
		fmt.Fscan(in, &m)
		cnts := make([]int, 30)
		var cur int
		mxkorob := 0
		for korob := 0; korob < m; korob++ {
			fmt.Fscan(in, &cur)
			cur++
			cur--
			cnts[cur]++
		}
		for i, v := range cnts {
			if v != 0 {
				mxkorob = i
			}
		}
		var curSpace, loads int
		for {
			stop := true
			for _, v := range cnts {
				if v > 0 {
					stop = false
				}
			}
			if stop {
				break
			}
			loads++
		CARLOOP:
			for car := 0; car < n; car++ {
				curSpace = k
				for i := mxkorob; i >= 0; i-- {
					for curSpace >= 1<<i {
						if cnts[i] == 0 {
							break
						} else {
							curSpace -= 1 << i
							cnts[i]--
						}
						if curSpace == 0 {
							continue CARLOOP
						}
					}
				}
			}
		}
		fmt.Fprintln(out, loads)
	}

}
