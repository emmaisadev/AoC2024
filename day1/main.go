package main

import (
	"github.com/ejfrick/advent-of-code-2024/utils"
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	exitCode := 0
	err := Execute()
	if err != nil {
		exitCode = 1
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(exitCode)
}

func Execute() error {
	input, err := utils.ReadInput()
	if err != nil {
		return err
	}

	updatedInput := strings.ReplaceAll(string(input), "   ", "\n")	

	scanner := bufio.NewScanner(strings.NewReader(updatedInput))
	var ints []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		ints = append(ints, i)
	}

	var leftList []int
	var rightList []int

	for i := 0; i < 2000; i = i + 2 {
		leftList = append(leftList, ints[i])
	}

	for i := 1; i < 2000; i = i + 2 {
		rightList = append(rightList, ints[i])
	}
	
	sort.Slice(leftList, func(a, j int) bool {
		return leftList[a] < leftList[j]
	})
	sort.Slice(rightList, func(a, j int) bool {
		return rightList[a] < rightList[j]
	})

	var distances []int

	for i := 0; i < 1000; i++ {
		if leftList[i] > rightList[i] {
			distances = append(distances, (leftList[i] - rightList[i]))
		}
		if leftList[i] < rightList[i] {
			distances = append(distances, (rightList[i] - leftList[i]))
		}
		if leftList[i] == rightList[i] {
			distances = append(distances, (leftList[i] - rightList[i]))
		}
	}

	totalDistance := 0

	for i := 0; i < 1000; i++ {
		totalDistance = totalDistance + distances[i]
	}

	similarityScore := 0

	for i := 0; i < 1000; i++ {
		numberofTimes := 0
		for o := 0; o < 1000; o++ {
			if leftList[i] == rightList[o] {
				numberofTimes++
			}
		}
		similarityScore = (similarityScore + (leftList[i] * numberofTimes)) 
	}

	fmt.Println(totalDistance)
	fmt.Println(similarityScore)
	return nil
}
