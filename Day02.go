package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("02/Day02-ross.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	safe := 0
	report_no := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		inc := parts[0] < parts[1] //signifies if increase or decrease
		report_no++

		for i := 0; i < len(parts)-1; i++ {
			v1, _ := strconv.Atoi(parts[i])
			v2, _ := strconv.Atoi(parts[i+1])
			differ := v1 - v2

			if inc && differ < 0 {
				differ = -differ
			} else if inc && differ > 0 {
				safe--
				fmt.Printf("%d must increase %d\n", report_no, differ)
				break
			} else if !inc && differ < 0 {
				safe--
				fmt.Printf("%d must decrease %d\n", report_no, differ)
				break
			}

			if !(differ >= 1 && differ <= 3) {
				safe--
				fmt.Printf("%d too much %d\n", report_no, differ)
				break
			}
			if i == len(parts)-2 {
				fmt.Printf("%d safe \n", report_no)
			}
		}
		safe++

	}

	fmt.Println(safe)
}
