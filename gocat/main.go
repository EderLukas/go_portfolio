package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("There must be at least one file name given as a parameter!")
		os.Exit(1)
	}

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			log.Printf("Error occured when opening file: %s\n%s", fname, err)
			f.Close()
			continue
		}
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Error on output %s\n%s", fname, err)
		}
		f.Close()
	}
}
