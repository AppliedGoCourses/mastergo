package main

import (
	"fmt"
	"github.com/AppliedGoCourses/mastergo/4_GoDevelopment/4-1-1_Compiling/gotools/procinfo"
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
