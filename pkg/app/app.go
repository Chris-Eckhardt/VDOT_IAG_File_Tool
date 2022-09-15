package app

import (
	"bufio"
	"fmt"

	//"log"
	"os"
	//"strings"
)

func Lookup(paramType string, param string, filePath string) {

	var matches []string

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// .. skip the header
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()

		if paramType == "id" {
			check := line[4:14]
			if check == param {
				matches = append(matches, line)
			}
		} else if paramType == "class" {
			check := line[80:84]
			if check == param {
				matches = append(matches, line)
			}
		} else if paramType == "protocol" {
			check := line[75:78]
			if check == param {
				matches = append(matches, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if len(matches) == 0 {
		fmt.Println("No matches found.")
	} else {
		printMatches(matches)
		fmt.Printf("%d matche(s) found. \n", len(matches))
	}
}

func printMatches(matches []string) {
	fmt.Println("--------------------------------------------")
	for _, tag := range matches {
		fmt.Printf("issuing agency: %s \n", tag[0:4])
		fmt.Printf("tag id: %s \n", tag[4:14])
		fmt.Printf("status: %c \n", tag[14])
		fmt.Printf("acct info: %s \n", tag[15:21])
		fmt.Printf("home agency: %s \n", tag[21:25])
		fmt.Printf("tag type: %c \n", tag[25])
		fmt.Printf("acct no: %s \n", tag[25:75])
		fmt.Printf("protocol: %s \n", tag[75:78])
		fmt.Printf("tag type: %c \n", tag[78])
		fmt.Printf("mount: %c \n", tag[79])
		fmt.Printf("IAG class: %s \n", tag[80:84])
		fmt.Println("--------------------------------------------")
	}
}
