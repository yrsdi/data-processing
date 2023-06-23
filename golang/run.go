package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("/home/yrsdi/Downloads/1000000-sales-records.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	startTime := time.Now()

	count := 0
	for scanner.Scan() && count < 1_000_000 {
		record := scanner.Text()
		// Process the record here
		fmt.Println(record)
		count++
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Elapsed time:", elapsedTime)
}
