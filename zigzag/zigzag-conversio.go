package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	_ = "PAHNAPLSIIGYIR"
	numRow1 := 3

	r := convert(s, numRow1)

	fmt.Println("resp : " + r)

}

func convert(s string, numRows int) (r string) {
	for i := range len(s) {
		if i%((numRows-1)*2) == 0 {
			r += string(s[i])
			continue
		}
	}
	for i := range len(s) {
		if i%2 == 1 {
			r += string(s[i])
			continue
		}
	}
	i := numRows - 1
	for i = range len(s) {
		if i%1 == 0 {
			r += string(s[i])
			continue
		}
	}

	return
}
