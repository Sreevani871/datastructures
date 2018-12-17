package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// logic
func arrLR(a []int32, d int32) []int32 {
	var i int32
	if len(a) < 2 {
		return a
	}
	for i = 0; i < d; i++ {
		temp := a[0]
		a = a[1:]
		a = append(a, temp)
	}
	return a
}

func main() {
	var (
		arr     []int32
		lrCount int
		arrLen  int
	)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 2; i++ {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		text = strings.Replace(text, "\n", "", -1)
		if i == 0 {
			splitText := strings.Split(text, " ")
			arrLen, _ = strconv.Atoi(splitText[0])
			lrCount, _ = strconv.Atoi(splitText[1])
		} else if i == 1 {
			arrStr := strings.Split(text, " ")
			if len(arrStr) != arrLen {
				fmt.Println("Invalid input")
				return
			}
			for _, e := range arrStr {
				eint, err := strconv.Atoi(e)
				if err != nil {
					fmt.Println(err)
					return
				}
				arr = append(arr, int32(eint))
			}
		}
	}
	fmt.Println("Arr:", arr)
	fmt.Println("LR: ", arrLR(arr, int32(lrCount)))

}
