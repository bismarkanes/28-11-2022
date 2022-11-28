package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func timeInWords(h int32, m int32) string {
	// Write your code here
	type MinuteModifier struct {
		// words of minute
		minuteWords string
		// relative hour to add
		hourToAdd int32
	}

	var NumberWords = map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		21: "twenty one",
		22: "twenty two",
		23: "twenty three",
		24: "twenty four",
		25: "twenty five",
		26: "twenty six",
		27: "twenty seven",
		28: "twenty eight",
		29: "twenty nine",
	}

	var MapMinutes = map[int32]MinuteModifier{
		15: {
			minuteWords: "quarter past",
			hourToAdd:   0,
		},
		30: {
			minuteWords: "half past",
			hourToAdd:   0,
		},
		45: {
			minuteWords: "quarter to",
			hourToAdd:   1,
		},
	}

	suffixwords := ""

	mapMin, ok := MapMinutes[m]

	if !ok {
		mapMin.minuteWords = ""

		strMin := "minutes"

		if m == 0 {
			suffixwords = "o' clock"
		} else if m < 30 && m > 0 {
			// before 30 mins relative to current hour
			if m == 1 {
				strMin = "minute"
			}
			mapMin.minuteWords = fmt.Sprintf("%s %s past", NumberWords[m], strMin)
		} else if m > 30 && m < 60 {
			// after 30 mins relative to next hour
			if (60 - m) == 1 {
				strMin = "minute"
			}
			mapMin.minuteWords = fmt.Sprintf("%s %s to", NumberWords[60-m], strMin)
			mapMin.hourToAdd = 1
		}
	}

	// hour to add relative
	h = h + mapMin.hourToAdd
	if h > 12 {
		// above 12 hour rotate
		h = h % 12
	}

	hourWords, ok := NumberWords[h]
	if !ok {
		return ""
	}

	// remove white space
	return strings.TrimSpace(fmt.Sprintf("%s %s %s", mapMin.minuteWords, hourWords, suffixwords))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

	fmt.Fprintf(writer, "%s\n", result)

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
