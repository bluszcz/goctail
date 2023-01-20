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
	LogPrintln("readReverseChunk", fmt.Sprint(start))

	// file.Seek(512, )
	data := make([]byte, BUFSIZ)

	// var offsetType int
	// if start == filesize {
	// 	offsetType = os.SEEK_END
	// } else {
	// 	LogPrintln(fmt.Sprint("Current"))
	// 	offsetType = os.SEEK_CUR
	// }

	// file.Seek(i6t64(BUFSIZ)*-1, offsetType)
	file.Seek(int64(start)-int64(BUFSIZ), os.SEEK_SET)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data[:count])

	return data[:count]
}

func processChunkedData(result string, lines int, start int, filesize int, file *os.File) {
	var amountEndlines = strings.Count(result, "\n")
	LogPrintln("How many endlines", fmt.Sprint(amountEndlines), fmt.Sprint(lines))
	if amountEndlines == lines {
		fmt.Printf("%s\n", result)
	} else if amountEndlines > lines+1 {
		for amountEndlines > lines {
			LogPrintln(fmt.Sprint("a"))
			amountEndlines = strings.Count(result, "\n")
			LogPrintln("How many endlines2", fmt.Sprint(amountEndlines), fmt.Sprint(lines))

			var indexN = strings.Index(result, "\n")
			amountEndlines = strings.Count(result, "\n")
			LogPrintln("How many endlines3", fmt.Sprint(amountEndlines), fmt.Sprint(lines))

			LogPrintln("indexX", fmt.Sprint(indexN))
			// if indexN > 0 {
			result = result[indexN+1:]
			// } else {
			// result = result[1:]
			// }
			LogPrintln(fmt.Sprint(amountEndlines))

			LogPrintln("amountEndLines", fmt.Sprint(amountEndlines))
		}
		fmt.Printf("%s", result)
	} else {
		LogPrintln(fmt.Sprint("Gotta do else"))
		var newResult string
		start = start - BUFSIZ
		newResult = fmt.Sprint(readReverseChunk(start, filesize, file))
		newResult = newResult + result
		processChunkedData(result, lines, start, filesize, file)
	}
}

func ReturnLastLines(lines int, filename string) {
	file, filesize := GetFile(filename)
	LogPrintln("filesize: " + fmt.Sprint(filesize))
	var result string = fmt.Sprintf("%s", readReverseChunk(int(filesize), int(filesize), file))
	processChunkedData(result, lines, int(filesize), int(filesize), file)
	// var amountEndlines = strings.Count(result, "\n")
	// LogPrintln("How many endlines", fmt.Sprint(amountEndlines))
	// // data := make([]byte, 100)
	// // count, err := file.Read(data)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // var indexN = strings.Index(result, "\n")
	// if amountEndlines == lines {
	// 	fmt.Printf("%s\n", result)
	// } else if amountEndlines > lines+1 {

	// 	for amountEndlines > lines {
	// 		LogPrintln(fmt.Sprint("a"))
	// 		var indexN = strings.Index(result, "\n")

	// 		LogPrintln("indexX", fmt.Sprint(indexN))
	// 		if indexN > 0 {
	// 			result = result[indexN:]
	// 		} else {
	// 			result = result[1:]
	// 		}
	// 		// fmt.Printf(fmt.Sprint(strings.Index(result, "\n")))

	// 		// indexN = strings.Index(result, "\n")
	// 		LogPrintln(fmt.Sprint(amountEndlines))

	// 		amountEndlines = strings.Count(result, "\n")
	// 		LogPrintln("amountEndLines", fmt.Sprint(amountEndlines))
	// 	}

	// 	fmt.Printf("%s", result)

	// }

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
