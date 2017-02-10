package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	for i := 0; i < len(os.Args); i++ {

		fmt.Printf("%d %s", i, os.Args[i] + "\r\n")
	}

	s := strings.Join(os.Args, " ");

	fmt.Printf(s)
}