package main

import (
	"bufio"
	"fmt"
	"os"
)

type chng struct {
	sec  int
	name string
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tests int
	var n, id, sec int
	var name, command string

	fmt.Fscan(in, &tests)
	//getnum := 0
	for test := 0; test < tests; test++ {
		ids := make(map[int][]chng)
		idtoprod := make(map[int]string)
		prods := make(map[string]int)
		fmt.Fscan(in, &n)
		for s := 1; s <= n; s++ {
			fmt.Fscan(in, &command)
			if command == "CHANGE" {
				fmt.Fscan(in, &name, &id)
				if oldid, ok := prods[name]; ok {
					ids[oldid] = append(ids[oldid], chng{sec: s, name: "404"})
					delete(idtoprod, oldid)
				}
				if oldprod, ok := idtoprod[id]; ok {
					delete(prods, oldprod)
				}
				ids[id] = append(ids[id], chng{sec: s, name: name})
				prods[name] = id
				idtoprod[id] = name
			} else {
				//getnum++
				fmt.Fscan(in, &id, &sec)
				if history, ok := ids[id]; !ok {
					fmt.Fprintln(out, 404)
				} else {
					prev := "404"
					ret := "404"
					for _, c := range history {
						if c.sec > sec {
							break
						}
						prev = c.name
					}
					ret = prev
					fmt.Fprintln(out, ret)
				}
			}

		}
	}
}
