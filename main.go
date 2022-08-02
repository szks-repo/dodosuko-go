package main

import (
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
	sb := &strings.Builder{}
	rand.Seed(time.Now().UnixNano())

	var (
		i           = 0
		numDodosuko = 0
	)
	for {
		if i == 100000000 {
			break
		}

		s := src[rand.Intn(2)]
		sb.WriteString(s)
		fmt.Print(s)
		if pass, dodosuko := dodosukoTest(sb.String()); !pass {
			sb.Reset()
			numDodosuko = 0
			fmt.Println("")
		} else if dodosuko {
			numDodosuko++
			if numDodosuko == 3 {
				fmt.Println("ラブ注入❤")
				break
			}
			sb.Reset()
		}
		i++
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
