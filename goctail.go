package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var flagvar int
var lastlinesvar int
var lastbytestvar int
var boolvar bool
var debug string

var BUFSIZ int = 512

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func LogPrintln(line ...string) {
	if debug == "true" {
		log.Println(line)
	}
}

// date time = 2023-01-18 16:18:22
// [1-2][0-9][0-9][0-9]-[0-1][0-9]-[0-3][0-9z]
// [0-2][0-9]:[0-5][0-9]:[0-5][0-9]
func processLine(line []byte) []byte {
	var regDate = regexp.MustCompile(`([1-2][0-9][0-9][0-9]\-[0-1][0-9]\-[0-3][0-9z])`)
	var regTime = regexp.MustCompile(`([0-9][0-9]\:[0-1][0-9]\:[0-3][0-9z])`)
	var returnLine = regDate.ReplaceAll(line, []byte(Blue+"$1"+Reset))
	var returnLine2 = regTime.ReplaceAll(returnLine, []byte(Purple+"$1"+Reset))
	return returnLine2
}

func GetFile(filename string) (*os.File, int64) {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var filesize = stats.Size()
	return file, filesize
}

// -c
func ReturnLastCount(count int, filename string) {
	file, filesize := GetFile(filename)
	LogPrintln("filesize: " + fmt.Sprint(filesize))

	file.Seek(int64(count)*-1, os.SEEK_END)
	data := make([]byte, count)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data[:count])
}

// Performed reversed read in BUFSIZ chunks
func readReverseChunk(start int, filesize int, file *os.File) string {
	LogPrintln("readReverseChunk", fmt.Sprint(start))
	var tmpBufSiz int
	if start >= BUFSIZ {
		tmpBufSiz = BUFSIZ
		file.Seek(int64(start)-int64(BUFSIZ), os.SEEK_SET)

	} else {
		tmpBufSiz = start
		start = 0
		file.Seek(0, 0)
	}
	data := make([]byte, tmpBufSiz)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", data[:count])
}

func processChunkedData(result string, lines int, start int, filesize int, file *os.File) {
	var amountEndlines = strings.Count(result, "\n")
	LogPrintln("How many endlines", fmt.Sprint(amountEndlines), fmt.Sprint(lines))
	if amountEndlines > lines {
		var indexN = strings.Index(result, "\n")
		result = result[indexN+1:]
		for amountEndlines > lines+1 {
			LogPrintln(fmt.Sprint("a"))
			amountEndlines = strings.Count(result, "\n")
			LogPrintln("How many endlines2", fmt.Sprint(amountEndlines), fmt.Sprint(lines))
			var indexN = strings.Index(result, "\n")
			amountEndlines = strings.Count(result, "\n")
			LogPrintln("How many endlines3", fmt.Sprint(amountEndlines), fmt.Sprint(lines))
			LogPrintln("indexX", fmt.Sprint(indexN))
			result = result[indexN+1:]
			LogPrintln(fmt.Sprint(amountEndlines))
			LogPrintln("amountEndLines", fmt.Sprint(amountEndlines))
		}
		// fmt.Printf("%s", result)
		fmt.Printf("%s", processLine([]byte(result)))
	} else {
		LogPrintln(fmt.Sprint("Gotta do else"))
		var newResult string
		start = start - BUFSIZ
		if start > 0 {
			newResult = fmt.Sprint(readReverseChunk(start, filesize, file))
			newResult = newResult + result
			processChunkedData(newResult, lines, start, filesize, file)
		} else {
			fmt.Printf("%s", processLine([]byte(result)))
		}
	}
}

// -n
func ReturnLastLines(lines int, filename string) {
	file, filesize := GetFile(filename)
	LogPrintln("filesize: " + fmt.Sprint(filesize))
	var result string = fmt.Sprintf("%s", readReverseChunk(int(filesize), int(filesize), file))
	processChunkedData(result, lines, int(filesize), int(filesize), file)
}

// https://stackoverflow.com/questions/35809252/check-if-flag-was-provided-in-go
// this flag module is poor
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	var line []byte = []byte("2023-01-18 16:18:21 configure openjdk-11-jre:amd64 11.0.17+8-1ubuntu2~22.04 <none>")

	processLine(line)
	debug = os.Getenv("DEBUG")
	// Need to set priority of parsing flags
	flag.IntVar(&lastlinesvar, "n", 10, "last lines to show")
	flag.IntVar(&lastbytestvar, "c", 100, "last bytes to show")
	flag.BoolVar(&boolvar, "f", false, "help message for flagname")

	flag.Parse()
	LogPrintln(fmt.Sprintf("c %t", isFlagPassed("c")))
	LogPrintln(fmt.Sprintf("n %t", isFlagPassed("n")))
	files := flag.Args()
	LogPrintln("files" + strings.Join(files, "-"))
	LogPrintln("last lines ", fmt.Sprint(lastlinesvar))
	LogPrintln("boolvar ", fmt.Sprint(boolvar))
	LogPrintln("Hello, World!")

	if isFlagPassed("c") {
		ReturnLastCount(lastbytestvar, files[0])
	} else {
		ReturnLastLines(lastlinesvar, files[0])
	}
}
