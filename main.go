package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	cidr    string
	outfile string
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Usage: masscidr -i cidr.txt -o outfile \n")
		os.Exit(0)
	}

	for i, arg := range os.Args {
		if arg == "-i" || arg == "--cidr" {
			cidr = os.Args[i+1]
		}
		if arg == "-o" || arg == "--output" {
			outfile = os.Args[i+1]
		}
	}

	cidrfile, error := os.Open(cidr)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(cidrfile)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	cidrfile.Close()

	out, err := os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	for _, each_ln := range text {

		_, err2 := out.WriteString("masscan -p1-65535 --rate=10000 " + each_ln + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	out.Close()

}
