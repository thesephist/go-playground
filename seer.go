package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const ANSI_BOLD string = "\033[1m"
const ANSI_END string = "\033[0m"
const ANSI_UNDER string = "\033[4m"
const ANSI_COLOR string = "\u001b[31m"

type row struct {
	startLineNumber int      // line number of the first line of this row in the source file
	values          []string // values of the row, divided into each column
}

func searchCSV(source io.Reader) (rows []row, err error) {
	curLineNumber := 1
	curRow := row{
		startLineNumber: 1,
		values:          []string{},
	}
	curStr := []byte{}
	inQuotedStr := false

	buf := make([]byte, 32)

	for {
		n, err := source.Read(buf)

		// parse lines
		for _, char := range buf[:n] {
			switch char {
			case '\\': // escape character
				continue
			case ',': // value separator
				if inQuotedStr {
					curStr = append(curStr, char)
				} else {
					curRow.values = append(curRow.values, string(curStr))
					curStr = []byte{}
				}
			case '\n': // new CSV rows
				if inQuotedStr {
					curStr = append(curStr, char)
				} else {
					// last value doesn't require a comma
					curRow.values = append(curRow.values, string(curStr))
					curStr = []byte{}

					rows = append(rows, curRow)
					curLineNumber++
					curRow = row{
						startLineNumber: curLineNumber,
						values:          []string{},
					}
				}
			case '"': // quoted values in CSV
				inQuotedStr = !inQuotedStr
			default:
				curStr = append(curStr, char)
			}
		}

		if err == io.EOF {
			err = nil
			// stop reading file
			break
		} else if err != nil {
			// return, preserve error
			return rows, err
		}
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (r row) String() string {
	result := strconv.Itoa(r.startLineNumber) + ":\t"
	for _, str := range r.values {
		result += str + "\n"
	}
	return result[:len(result)-1] // remove trailing comma
}

func (r row) Content() string {
	return strings.Join(r.values, " ")
}

func (r row) Match(keyword string) bool {
	return strings.Contains(r.Content(), keyword)
}

func (r row) Matches(keywords []string) bool {
	for _, str := range keywords {
		if r.Match(str) {
			return true
		}
	}
	return false
}

func colorSubstring(s, sub string) string {
	return strings.Replace(s, sub, ANSI_COLOR + ANSI_BOLD +sub+ANSI_END, -1)
}

func colorSubstrings(s string, subs []string) string {
	result := s
	for _, sub := range subs {
		result = colorSubstring(result, sub)
	}
	return result
}

// Searches a CSV for a given string, and shows those rows (not file lines)
func main() {
	fileName := os.Args[1]
	args := os.Args[2:]

	// open file
	fileReader, err := os.Open(fileName)
	rows, err := searchCSV(fileReader)

	fmt.Println(ANSI_UNDER + "Matched rows:" + ANSI_END)
	for _, r := range rows {
		if r.Matches(args) {
			fmt.Println(colorSubstrings(r.String(), args))
		}
	}

	check(err)
}
