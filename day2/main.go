package main

import (
    "github.com/ejfrick/advent-of-code-2024/utils"
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
    //"sort"
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
	
	var reports [][]int

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {
		var report []int
		// general idea: operate on one report at a time, generate a "report" slice object for this report, which is then added to the larger "reports" slice 
		reportScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		reportScanner.Split(bufio.ScanWords)
		for reportScanner.Scan() {
			i, err := strconv.Atoi(reportScanner.Text())
			if err != nil {
				return err
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	var safeReports [][]int

	for _, report := range reports {
		var reportSafe bool = true
		var decreasing bool = false
		var increasing bool = false
		var previousValue int = 0
		var numOfProblems int = 0
		for i, value := range report {
			if i == 0 {
				previousValue = value
				continue
			}
			switch {
				case value == previousValue:
					reportSafe = false
					numOfProblems++
					previousValue = value
				case value < (previousValue - 3):
					reportSafe = false
					numOfProblems++
					previousValue = value
				case value > (previousValue + 3):
					reportSafe = false
					numOfProblems++
					previousValue = value
				case value < previousValue:
					decreasing = true
				case value > previousValue:
					increasing = true
			}
			previousValue = value
			if increasing && decreasing {
				reportSafe = false
				numOfProblems++
			}
			if !reportSafe && numOfProblems > 1{
				// you need to actually remove the single problem element from the report, one at a time, and attempt to run the report again and see if it passes. make copies of each report with the bad element removed instead.
				continue
			}
			if !reportSafe && numOfProblems == 1 {
				fmt.Println(report)
			}
			reportSafe = true
		}
		if reportSafe {
			safeReports = append(safeReports, report)
		}
	}
	fmt.Println(safeReports)
	fmt.Println(len(safeReports))

	return nil
}
