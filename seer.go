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
const ANSI_COLOR string = "\u001b[31m" // ANSI red escape

type row struct {
	header          []string // pointer to the header row string slice
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
				curLineNumber++
				if inQuotedStr {
					curStr = append(curStr, char)
				} else {
					// last value doesn't require a comma
					curRow.values = append(curRow.values, string(curStr))
					curStr = []byte{}

					rows = append(rows, curRow)
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
	headers, rows := rows[0].values, rows[1:]
	for idx := range rows {
		rows[idx].header = headers
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clip(s string, max int) string {
	if max > len(s) {
		return s
	} else {
		return s[:max]
	}
}

func (r row) String() string {
	result := ANSI_UNDER + "Line " + strconv.Itoa(r.startLineNumber) + ":\n" + ANSI_END
	for idx, str := range r.values {
		if str != "" {
			result += "    " + ANSI_BOLD + clip(r.header[idx], 10) + ANSI_END + ":\t" + str + "\n"
		}
	}
	return result[:len(result)-1]
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
	return strings.Replace(s, sub, ANSI_COLOR+ANSI_BOLD+sub+ANSI_END, -1)
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: seer [CSV file to search] [... search keywords]")
		os.Exit(1)
	}

	fileName := os.Args[1]
	args := os.Args[2:]

	fileReader, err := os.Open(fileName)
	rows, err := searchCSV(fileReader)

	matchedRows := []row{}
	for _, r := range rows {
		if r.Matches(args) {
			matchedRows = append(matchedRows, r)
		}
	}

	fmt.Println(ANSI_UNDER + strconv.Itoa(len(matchedRows)) + " Matched rows:\n" + ANSI_END)
	for _, r := range matchedRows {
		fmt.Println(colorSubstrings(r.String(), args))
	}
	check(err)
}
