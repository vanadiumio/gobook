package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("You need to provide the filename")
		os.Exit(1)
	}
	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		var wordcount int
		for scanner.Scan() {
			wordcount++
		}
		if scanner.Err() != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Found", wordcount, "words")
	}
}
