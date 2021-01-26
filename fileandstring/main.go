package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getagentdetails(filename string) (agentname string, agentversion string) {

	file, err := os.Open(filename)
	if err != nil {
		return
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	if fileScanner.Err() != nil {
		fmt.Printf(" > Failed with error %v\n", fileScanner.Err())
		return
	}
	lineCount := 0
	var firstline string
	for fileScanner.Scan() {
		lineCount++
		if lineCount == 1 { //Check firstline only
			firstline = fileScanner.Text()

		}

	}

	if lineCount == 0 { //return null value if file is empty or have two lines

		fmt.Println("Invalid file : File is empty  ")
		return agentname, agentversion
	}

	//check string

	wordList := strings.Fields(firstline)
	if len(wordList) != 1 {

		fmt.Println("invalid file : Found multiple details")
		return agentname, agentversion

	}
	//parsing data
	AgentName := strings.Split(firstline, "-")[0]   //parsing name
	AgentName = strings.TrimSpace(AgentName)        // trimming space
	AgentVer := strings.Trim(firstline, AgentName)  //parsing version details
	AgentVersion := strings.TrimLeft(AgentVer, "-") // trimmimg
	AgentVersion = strings.TrimSpace(AgentVersion)  // trimming space

	agentname = AgentName
	agentversion = AgentVersion

	return agentname, agentversion

}

func main() {

	a, b := getagentdetails("a.txt")
	fmt.Printf("name=%s version=%s", a, b)
}
