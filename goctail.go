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
var boolvar bool
var debug string

func LogPrintln(line string) {
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
	fmt.Printf("The file is %d bytes long", filesize)
	return file, filesize
}

func ReturnLastCount(count int, filename string) {
	file, filesize := GetFile(filename)
	fmt.Printf("The file is %d bytes long\n", filesize)
	file.Seek(int64(count)*-1, os.SEEK_END)
	data := make([]byte, count)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: \n%s\n", count, data[:count])
}

func ReturnLastLines(lines int, filename string) {
	file, filesize := GetFile(filename)
	fmt.Printf("The file is %d bytes long", filesize)
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

func main() {
	var debug string

	flag.IntVar(&lastlinesvar, "n", 10, "last lines to show")
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.BoolVar(&boolvar, "f", false, "help message for flagname")

	flag.Parse()
	files := flag.Args()
	LogPrintln("files" + strings.Join(files, "-"))

	if debug == "true" {
		fmt.Println("flagvar has value ", flagvar)
		fmt.Println("last lines ", lastlinesvar)
		fmt.Println("boolvar ", boolvar)
		fmt.Println("Hello, World!")
		fmt.Println(" DEBUG <<<<")
		fmt.Println("")

	}
	ReturnLastCount(lastlinesvar, files[0])
	// ReturnLastLines(lastlinesvar, files[0])
}
