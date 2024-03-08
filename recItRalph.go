package main

import (
	"encoding/csv"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
)

func banner() string {
	recBanner, err := os.ReadFile("./docs/banner.txt")
	if err != nil {
		log.Fatal("Bad banner read", err)
	}
	fmt.Println(string(recBanner))
	return (string(recBanner))
}

func readData(filePath string) [][]string {

}

func main() {

	banner()

	parser := argparse.NewParser("commands", "simple ags for recs")

	qbPath := parser.String("q", "qbPath", &argparse.Options{Required: true, Help: "The path to the qb.csv", Default: "null"})

	bkPath := parser.String("b", "bkPath", &argparse.Options{Required: true, Help: "The path to the bank.csv", Default: "null"})

	format := parser.String("f", "format", &argparse.Options{Required: false, Help: "Choose output format: csv json shell", Default: "shell"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	fmt.Printf("q: %s\n", *qbPath)
	fmt.Printf("b: %s\n", *bkPath)
	fmt.Printf("f: %s\n", *format)

	bkFile, err := os.Open(*bkPath)

	if err != nil {
		log.Fatal("Error while opening bank file", err)

	}

	defer bkFile.Close()

	bkReader := csv.NewReader(bkFile)

	bkRecords, err := bkReader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading bank file", err)

	}

	for _, eachRecord := range bkRecords {
		fmt.Println(eachRecord)
	}
}
