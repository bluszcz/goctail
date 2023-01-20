package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var flagvar int
var lastlinesvar int
var lastbytestvar int
var boolvar bool
var debug string

var BUFSIZ int = 512

func LogPrintln(line ...string) {
	if debug == "true" {
		log.Println(line)
	}
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
	// fmt.Printf("The file is %d bytes long", filesize)
	return file, filesize
}

func ReturnLastCount(count int, filename string) {
	file, filesize := GetFile(filename)
	LogPrintln("filesize: " + fmt.Sprint(filesize))

	file.Seek(int64(count)*-1, os.SEEK_END)
	data := make([]byte, count)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data[:count])
}

// Performed reversed read in BUFSIZ chunks
func readReverseChunk(start int, filesize int, file *os.File) []byte {
	LogPrintln("readReverseChunk")

	// file.Seek(512, )
	data := make([]byte, BUFSIZ)

	var offsetType int
	if start == filesize {
		offsetType = os.SEEK_END
	} else {
		offsetType = os.SEEK_CUR
	}

	file.Seek(int64(BUFSIZ)*-1, offsetType)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", data[:count])

	return data[:count]
}

func ReturnLastLines(lines int, filename string) {
	file, filesize := GetFile(filename)
	LogPrintln("filesize: " + fmt.Sprint(filesize))
	var result string = fmt.Sprintf("%s", readReverseChunk(int(filesize), int(filesize), file))
	var amountEndlines = strings.Count(result, "\n")
	LogPrintln("How many endlines", fmt.Sprint(amountEndlines))
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var indexN = strings.Index(result, "\n")
	if amountEndlines == lines {
		fmt.Printf("%s\n", result)
	} else if amountEndlines > lines {

		for amountEndlines > lines {
			// fmt.Printf(fmt.Sprint(strings.Index(result, "\n")))
			result = result[indexN:]
			amountEndlines = strings.Count(result, "\n")
		}
		fmt.Printf("%s", result)

	}

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
	debug = os.Getenv("DEBUG")
	// Need to set priority of parsing flags

	flag.IntVar(&lastlinesvar, "n", 10, "last lines to show")
	flag.IntVar(&lastbytestvar, "c", 100, "last bytes to show")
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.BoolVar(&boolvar, "f", false, "help message for flagname")

	flag.Parse()
	LogPrintln(fmt.Sprintf("c %t", isFlagPassed("c")))
	LogPrintln(fmt.Sprintf("n %t", isFlagPassed("n")))
	files := flag.Args()
	LogPrintln("files" + strings.Join(files, "-"))
	LogPrintln("flagvar has value ", fmt.Sprint(flagvar))
	LogPrintln("last lines ", fmt.Sprint(lastlinesvar))
	LogPrintln("boolvar ", fmt.Sprint(boolvar))
	LogPrintln("Hello, World!")

	if isFlagPassed("c") {
		ReturnLastCount(lastbytestvar, files[0])
	} else {
		ReturnLastLines(lastlinesvar, files[0])
	}
}
