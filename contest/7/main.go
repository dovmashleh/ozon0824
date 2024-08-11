package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type dkey struct {
	typ  byte
	diff int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tests int
	var k, n, m int

	fmt.Fscan(in, &tests)
TST:
	for test := 0; test < tests; test++ {
		fmt.Fscan(in, &k, &n, &m)
		rows := make([][]byte, n)
		cols := make([][]byte, m)
		diags := make(map[dkey][]byte)
		_, _ = in.ReadString('\n')
		for i := 0; i < n; i++ {
			str, _ := in.ReadString('\n')
			rows[i] = []byte(strings.Trim(str, " \r\n"))
			for j, v := range rows[i] {
				cols[j] = append(cols[j], v)
				d1key := dkey{'g', i - j}
				d2key := dkey{'p', i + j}
				diags[d1key] = append(diags[d1key], v)
				diags[d2key] = append(diags[d2key], v)
			}
		}
		canwin := false

		if n < m {
			for _, line := range rows {
				can, won := checkForWin(line, k)
				if won {
					fmt.Fprintln(out, "NO")
					continue TST
				}
				canwin = canwin || can
			}
			for _, line := range cols {
				can, won := checkForWin(line, k)
				if won {
					fmt.Fprintln(out, "NO")
					continue TST
				}
				canwin = canwin || can
			}
		} else {
			for _, line := range cols {
				can, won := checkForWin(line, k)
				if won {
					fmt.Fprintln(out, "NO")
					continue TST
				}
				canwin = canwin || can
			}
			for _, line := range rows {
				can, won := checkForWin(line, k)
				if won {
					fmt.Fprintln(out, "NO")
					continue TST
				}
				canwin = canwin || can
			}
		}
		for _, line := range diags {
			can, won := checkForWin(line, k)
			if won {
				fmt.Fprintln(out, "NO")
				continue TST
			}
			canwin = canwin || can
		}

		if canwin {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
func checkForWin(line []byte, k int) (bool, bool) {
	lastdot := -1
	ln := 0
	dots := 0
	oln := 0
	conk := 0
	for i, v := range line {
		if v == 'X' {
			ln++
			if ln >= k {
				return true, false
			}
			oln = 0
			conk++
			if conk == k {
				return true, true
			}
		} else if v == 'O' {
			ln = 0
			dots = 0
			oln++
			if oln == k {
				return false, true
			}
			conk = 0
		} else if v == '.' {
			oln = 0
			if dots == 0 {
				lastdot = i
				dots++
				ln++
			} else {
				ln = i - lastdot
				lastdot = i

			}
			conk = 0
			if ln >= k {
				return true, false
			}
		}
	}

	return false, false
}
