package main

import (
	"fmt"
	"gotools/procinfo"
	"log"
	"os"
)

func main() {
	err := helper(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I am", os.Args[0])
	fmt.Println(procinfo.Get())
}
