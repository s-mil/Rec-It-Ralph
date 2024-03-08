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
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}


func main() {

	banner()

	parser := argparse.NewParser("commands", "simple ags for recs")

	qbPath := parser.String("q", "qbPath", &argparse.Options{Required: true, Help: "The path to the qb.csv", Default: "null"})

	bkPath := parser.String("b", "bkPath", &argparse.Options{Required: true, Help: "The path to the bank.csv", Default: "null"})

	format := parser.String("f", "format", &argparse.Options{Required: false, Help: "Choose output format: csv json shell", Default: "shell"})
	
	dflag := parser.Flag("d", "debug", &argparse.Options{Required: false, Help: "Run Debug options", Default: false})
	
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}



	
	bkRecords := readData(*bkPath)
	qbRecords := readData(*qbPath)


	if 	*dflag{
		fmt.Printf("q: %s\n", *qbPath)
		fmt.Printf("b: %s\n", *bkPath)
		fmt.Printf("f: %s\n", *format)

		for _, eachRecord := range bkRecords {
			fmt.Println(eachRecord)
		}

		fmt.Println("-------------------------------------------------------")

		for _, eachRecord := range qbRecords {
			fmt.Println(eachRecord)
		}
	}
}
