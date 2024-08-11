package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
TSTLOOP:
	for test := 0; test < tests; test++ {
		cnts := make(map[int]int)
		fmt.Fscan(in, &n)
		var num int
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			cnts[num]++
		}
		outp, _ := in.ReadString('\n')
		outp, _ = in.ReadString('\n')
		if outp[0] == ' ' {
			fmt.Fprintln(out, "NO")
			continue TSTLOOP
		}
		if outp[len(outp)-1] == ' ' {
			fmt.Fprintln(out, "NO")
			continue TSTLOOP
		}
		if strings.Contains(outp, "  ") {
			fmt.Fprintln(out, "NO")
			continue TSTLOOP
		}
		outar := strings.Fields(outp)
		if len(outar) != n {
			fmt.Fprintln(out, "NO")
			continue TSTLOOP
		}
		prev := math.MinInt
		for _, v := range outar {
			if v[0] == '0' ||
				len(v) > 1 && v[0] == '-' && v[1] == '0' {
				fmt.Fprintln(out, "NO")
				continue TSTLOOP
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintln(out, "NO")
				continue TSTLOOP
			}
			if num < prev {
				fmt.Fprintln(out, "NO")
				continue TSTLOOP
			}
			cnts[num]--
			if cnts[num] < 0 {
				fmt.Fprintln(out, "NO")
				continue TSTLOOP
			}
			prev = num
		}

		fmt.Fprintln(out, "YES")
	}
}
