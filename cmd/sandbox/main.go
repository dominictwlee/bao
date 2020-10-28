package main

import (
	"fmt"
	"os"
)

func main() {
	srcStat, err := os.Stat("dumb.txt")
	if err != nil {
		return
	}

	fmt.Println(srcStat.Mode().IsRegular())

}
