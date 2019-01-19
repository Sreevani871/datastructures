package main

import (
	"fmt"
	"math"
	"strings"
)

var (
	allowedCharsInUrl = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func idToShortUrl(id int) (url string) {
	var digits []int
	for id > 0 {
		digits = append(digits, id%62)
		id = id / 62
	}
	for _, digit := range digits {
		url = url + string(allowedCharsInUrl[digit])
	}
	return url

}

func shortUrlToId(url string) int {
	var (
		digits []int
		id     int
	)

	for i := 0; i < len(url); i++ {
		digits = append(digits, strings.Index(allowedCharsInUrl, string(url[i])))
	}

	j := 0
	for _, d := range digits {
		id = id + d*int(math.Pow(62, float64(j)))
		j++
	}
	return id
}

func main() {
	for id := 123456; id < (123456 + 10); id++ {
		str := idToShortUrl(id)
		// fmt.Printf("Id %d to string \"%s\"\n", id, str)
		idr := shortUrlToId(str)
		// fmt.Printf("string \"%s\" to id %d\n", str, idr)
		if id != idr {
			fmt.Println("Failed: ", id)
		}
	}
}
