package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

// serialize slices into string
func slicesToString(s [][]int32) string {
	result := ""
	for i := range s {
		for j := range s[i] {
			sn := strconv.Itoa(int(s[i][j]))
			result = result + sn
		}
	}
	return result
}

// calculate cost by comparing input square string and reference magic square string
func calculateCostMagicSquareReference(inputSquare string) int32 {
	var cost int32 = math.MaxInt32

	// reference for 3x3 magic squares
	magicSquares := []string{"618753294", "816357492", "834159672", "438951276", "672159834", "276951438", "294753618", "492357816"}

	for _, magicSquare := range magicSquares {
		var curCost int32 = calculateCost(inputSquare, magicSquare)

		if curCost < cost {
			cost = curCost
		}
	}

	return cost
}

func calculateCost(square, magicSquare string) int32 {
	var totalCost int32 = 0

	for i := range magicSquare {
		if magicSquare[i] != square[i] {
			n1, err := strconv.Atoi(string(magicSquare[i]))

			if err != nil {
				continue
			}

			n2, err := strconv.Atoi(string(square[i]))

			if err != nil {
				continue
			}

			var cost int32 = int32(n1 - n2)
			if cost < 0 {
				cost = -cost
			}

			totalCost = totalCost + cost
		}
	}

	return totalCost
}

func formingMagicSquare(inputSquare [][]int32) int32 {
	// Write your code here

	// result of total cost
	var totalCost int32 = 0

	// serialize input magic square number into string
	sInputSquare := slicesToString(inputSquare)

	// find matched reference magic square
	totalCost = calculateCostMagicSquareReference(sInputSquare)

	return totalCost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
