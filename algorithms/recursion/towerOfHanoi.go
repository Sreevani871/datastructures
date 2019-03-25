package main

import (
	"fmt"
)

func TowerOfHanoi(disk int, source, dest, aux rune) {
	if disk == 1 {
		fmt.Printf("Move disk:%d from source:%c to dest:%c\n", disk, source, dest)
		return
	}
	TowerOfHanoi(disk-1, source, aux, dest)
	fmt.Printf("Move disk:%d from source:%c to dest:%c\n", disk, source, dest)
	TowerOfHanoi(disk-1, aux, dest, source)
}

func main() {
	TowerOfHanoi(3, 'A', 'B', 'C')
}
