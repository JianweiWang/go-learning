package main

import (
	"bufio"
	"os"
	"fmt"
)

func Du(reader bufio.Reader)  {
	input := bufio.NewScanner(os.Stdin)
	countMap := make(map[string]int)

	for input.Scan() {
		countMap[input.Text()]++
	}

	for s, count := range(countMap) {
		fmt.Printf("%s: %d", s, count)
	}
}
