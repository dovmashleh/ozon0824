package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	//f, _ := os.OpenFile("in.txt", os.O_RDONLY, 0644)
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	fmt.Fscan(in, &n)
	summary := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		var cnt int
		var str []byte
		var total []byte
		fmt.Fscan(in, &cnt)
		for j := 0; j <= cnt; j++ {
			str, _, _ = in.ReadLine()
			total = append(total, str...)
		}
		a := make(map[string]any)
		b := make([]interface{}, 0, 0)
		json.Unmarshal([]byte(total), &a)
		if len(a) == 0 {
			json.Unmarshal([]byte(total), &b)
			b = removeEmpty(b)
			//fmt.Println(len(b), b)
			if len(b) != 0 {
				summary = append(summary, b)
			}
		} else {
			a = process(a)
			//fmt.Println(len(a), a)
			if len(a) != 0 {
				summary = append(summary, a)
			}
		}

	}
	result, _ := json.Marshal(summary)
	fmt.Fprintln(out, string(result))
}
func process(got map[string]any) map[string]any {
	for k, val := range got {
		//fmt.Printf("%s is %T \n", k, val)
		if sl, isSlice := val.([]interface{}); isSlice {
			sl = removeEmpty(sl)
			if len(sl) == 0 {
				delete(got, k)
			} else {
				got[k] = sl
			}
		}
		if mp, isMap := val.(map[string]any); isMap {
			processed := process(mp)
			if len(processed) == 0 {
				delete(got, k)
			} else {
				got[k] = processed
			}
		}
	}
	return got
}

func removeEmpty(in []interface{}) []interface{} {
	newSlice := make([]interface{}, 0, len(in))
	for _, v := range in {
		switch v.(type) {
		case string:
			if v.(string) != "" {
				newSlice = append(newSlice, v)
			}
		case map[string]interface{}:
			t := v.(map[string]interface{})
			t = process(t)
			if len(t) == 0 {
				continue
			} else {
				newSlice = append(newSlice, t)
			}
		case []interface{}:
			t := v.([]interface{})
			t = removeEmpty(t)
			if len(t) == 0 {
				continue
			} else {
				newSlice = append(newSlice, t)
			}

		}
	}
	return newSlice
}
