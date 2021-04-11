package main

import (
	"fmt"
	"os"
)

const (
	nloop = 10000
)

func main() {

	for i := 0; i < nloop; i++ {
		ppid := os.Getppid()
		fmt.Println(ppid)
	}

}
