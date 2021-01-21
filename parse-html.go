package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

		if name == "" || id == "" {
			continue
		}

		nameMap[id] = name
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Reading file failed: %v", err)
	}

	return nameMap
}

func writeVideoNameAndIdtoFile(videoNames map[string]string, filename string, folderToCheck string) {
	file, err := os.Create(filename)

	fileNames := make(map[string]bool, 0)

	if err != nil {
		log.Fatalf("Error occured creating new file: %v", err)
		return
	}

	fileInfo, err := ioutil.ReadDir(folderToCheck)

	if err != nil {
		log.Fatalf("Error occured reading video directory: %v", err)
		return
	}

	for _, file := range fileInfo {
		fileNames[file.Name()] = true
	}

	writer := bufio.NewWriter(file)

	for id, name := range videoNames {
		if fileNames[id] {
			written, err := writer.WriteString(fmt.Sprintf("%v: %v\n", id, name))

			if err != nil {
				log.Printf("Error occured while writing (written: %v): %v", written, err)
				return
			}

		}
	}

	err = writer.Flush()

	if err != nil {
		log.Printf("Error occured while flushing buffer: %v", err)
		return
	}
}

func main() {

	file, err := os.Open("names.html")

	if err != nil {
		log.Fatalf("Failed to open html file: %v", err)
		return
	}

	videoNames := ParseFile(file)

	writeVideoNameAndIdtoFile(videoNames, "1-shion-archive-stream-names-to-ids.txt", "../shion-435-archive")
}
