package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("Day01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var col1, col2 []int
	frequency := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		v1, _ := strconv.Atoi(parts[0])
		v2, _ := strconv.Atoi(parts[1])

		col1 = append(col1, v1)
		col2 = append(col2, v2)

		//in preparation for part 2
		frequency[v2] = frequency[v2] + 1

	}
	fmt.Println(col1)
	fmt.Println(col2)

	//sort slices
	sort.Ints(col1)
	sort.Ints(col2)

	fmt.Println(col1)
	fmt.Println(col2)

	total_dist := 0
	similarity_score := 0 //part2
	for i := 0; i < len(col1); i++ {
		distance := col1[i] - col2[i]
		if distance < 0 {
			distance = -distance
		}
		total_dist += distance
		similarity_score += frequency[col1[i]] * col1[i] //part2
	}

	fmt.Println("Day01 Part1: ", total_dist)

	//part #2
	fmt.Println("Day01 Part2: ", similarity_score)

}
