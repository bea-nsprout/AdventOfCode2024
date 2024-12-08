package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day02.txt")
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

		v1, _ := strconv.Atoi(parts[0])
		v2, _ := strconv.Atoi(parts[1])
		inc := v1 < v2 //signifies if increase or decrease
		//if inc {
		//	fmt.Print("INC: ")
		//} else {
		//	fmt.Print("DEC: ")
		//}

		report_no++

		for i := 0; i < len(parts)-1; i++ {
			v1, _ := strconv.Atoi(parts[i])
			v2, _ := strconv.Atoi(parts[i+1])
			differ := v1 - v2
			//fmt.Print(inc, differ)

			if inc && differ < 0 {
				differ = -differ
				//fmt.Printf("%d is correct (%d -> %d) %d\n", report_no, v1, v2, differ)
			} else if inc && differ > 0 {
				safe--
				//fmt.Printf("%d must increase (%d -> %d) %d\n", report_no, v1, v2, differ)
				break
			} else if !inc && differ < 0 {
				safe--
				//fmt.Printf("%d must decrease (%d -> %d) %d\n", report_no, v1, v2, differ)
				break
			}

			if !(differ >= 1 && differ <= 3) {
				safe--
				//fmt.Printf("%d too much (%d -> %d) %d\n", report_no, v1, v2, differ)
				break
			}

			//if i == len(parts)-2 {
			//	fmt.Printf("%d safe \n", report_no)
			//}
		}
		safe++

	}

	fmt.Println("total safe: ", safe)
}
