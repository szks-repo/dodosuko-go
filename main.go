package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var tests = []string{
	"ドド",
	"ドドスコ",
	"ドドスコスコ",
	"ドドスコスコスコ",
}

func main() {
	src := []string{"ドド", "スコ"}
	buf := &bytes.Buffer{}
	rand.Seed(time.Now().UnixNano())

	var numDodosuko int

	for {
		s := src[rand.Intn(2)]
		buf.WriteString(s)
		fmt.Print(s)
		if pass, dodosuko := dodosukoTest(buf.String()); !pass {
			buf.Reset()
			numDodosuko = 0
			fmt.Println("")
		} else if dodosuko {
			numDodosuko++
			if numDodosuko == 3 {
				fmt.Println("ラブ注入❤")
				return
			}
			buf.Reset()
		}
	}
}

func dodosukoTest(s string) (bool, bool) {
	for i := 0; i < len(tests); i++ {
		if !strings.HasPrefix(s, tests[i]) {
			return false, false
		}
		if s == tests[i] {
			if i != 3 {
				break
			}
			return true, true
		}
	}
	return true, false
}
