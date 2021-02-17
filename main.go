package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	ReadTxt("jobs.txt")
}
func ReadTxt(fileName string) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
