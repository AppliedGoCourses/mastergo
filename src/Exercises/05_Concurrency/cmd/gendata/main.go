package main

import (
	"fmt"
	"log"
	"os"

	gf "github.com/brianvoe/gofakeit"
)

func main() {
	f, err := os.Create("trainingdata.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		fmt.Fprintf(f, "\"%s\", ", gf.Name())
		go func() {
			for j := 0; j < 1000; j++ {
				fmt.Fprintf(f, "\"%d\", ", gf.Number(80, 160))
			}
		}()
	}
}
