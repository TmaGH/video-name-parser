package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseLine(line string) (name string, id string) {

	substrings := strings.Split(line, "\"")

	for i := range substrings {

		if strings.Contains(substrings[i], "title") && strings.Contains(substrings[i+1], ".mp4") {
			id = substrings[i+1]
			break
		}

	}

	substrings = strings.Split(line, ">")

	for i := range substrings {

		if strings.Contains(substrings[i], ".mp4") {
			name = strings.Split(substrings[i+1], "<")[0]
		}

	}

	return
}

func ParseFile(file *os.File) map[string]string {
	scanner := bufio.NewScanner(file)

	nameMap := make(map[string]string, 0)

	for scanner.Scan() {
		name, id := ParseLine(scanner.Text())

		nameMap[id] = name
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading file failed: %v", err)
	}

	for id, name := range nameMap {
		log.Printf("Id: %v, Name: %v", id, name)
	}

	return nameMap
}

func writeVideoNameAndIdtoFile(videoNames map[string]string, filename string, folderToCheck ) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Error occured creating new file: %v", err)
		return
	}

	writer := bufio.NewWriter(file)

	for id, name := range videoNames {
		writer.WriteString(fmt.Sprintf("%v: %v\n", id, name))
	}
}

func main() {

	file, err := os.Open("names.html")

	if err != nil {
		log.Fatalf("Failed to open html file: %v", err)
		return
	}

	videoNames := ParseFile(file)

	writeVideoNameAndIdtoFile(videoNames, "1-shion-archive-stream-names-to-ids.txt")
}
