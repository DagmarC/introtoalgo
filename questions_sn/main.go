package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/DagmarC/introtoalgo/questionsSN/page"
)

var o = flag.String("o", "", "file to be created to download data to, if empty os.Stdout will be used")

func main() {
	flag.Parse()

	if len(os.Args) <= 1 {
		fmt.Println(errors.New("not enough arguments, supply url and optional -o <filename> "))
	}
	url := os.Args[len(os.Args)-1]

	fmt.Println(*o)
	if *o == "" {
		fmt.Println("INFO: Writing to os.Stdout")
		err := page.DownloadPage(url, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("INFO: Creating and writing to ", *o)
	file, err := os.Create(*o)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	err = page.DownloadPage(url, file)
	if err != nil {
		fmt.Println(err)
	}
}
