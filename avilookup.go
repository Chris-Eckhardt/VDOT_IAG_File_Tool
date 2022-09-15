package main

import (
	"flag"
	"fmt"

	"github.com/chris-eckhardt/golang/pkg/app"
)

var (
	paramType *string
	param     *string
	filePath  *string
)

func init() {
	paramType = flag.String("type", "id", "Search parameter type.")
	param = flag.String("param", "", "Search parameter.")
	filePath = flag.String("path", "", "Path to IAG tag file")
}

func main() {
	flag.Parse()

	var pass bool = false

	if *paramType == "id" || *paramType == "class" || *paramType == "protocol" {
		pass = true
	} else {
		fmt.Println("Unknown parameter type.")
	}

	if *param == "" {
		fmt.Println("Parameter cant be null or empty.")
		pass = false
	}

	if *paramType == "id" && len(*param) != 10 {
		fmt.Println("id must have a 10 character parameter.")
		pass = false
	}

	if *paramType == "class" && len(*param) != 4 {
		fmt.Println("class must have a 4 character parameter.")
		pass = false
	}

	if *paramType == "protocol" && len(*param) != 3 {
		fmt.Println("protocol must have a 3 character parameter.")
		pass = false
	}

	if *filePath == "" {
		fmt.Println("File path cant be null or empty.")
		pass = false
	}

	if pass {
		app.Lookup(*paramType, *param, *filePath)
	}
}
