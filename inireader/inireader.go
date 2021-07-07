package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(filename, expectSection, expectKey string) string  {
	// open the file
	file, err := os.Open(filename)
	// can't find the file, return null
	if err != nil{
		return ""
	}
	// close the file when the function running over
	defer file.Close()

	// read the file via reader
	reader := bufio.NewReader(file)

	// get the name of the current section
	var sectionName string
	for  {
		// read one line of the file
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// trim the blank char at font and back size of the line string
		linestr = strings.TrimSpace(linestr)
		// ignore the blank line (no contents at this line)
		if linestr == ""{
			continue
		}
		// ignore comments
		if linestr[0] == ';'{
			continue
		}
		// get the section, key, and value
		// get the section, if the first and list char is [ and ], then it's the section line
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']'{
			// get the section name
			sectionName = linestr[1:len(linestr)-1]
		}else if sectionName == expectSection {
			// Split the key and value connected by '='
			pair := strings.Split(linestr, "=")
			// Gurantee what the pair we splited only contains one '='
			if len(pair) == 2{
				// remove extra blank char of key
				key := strings.TrimSpace(pair[0])
				// ensure the key is expected key
				if key == expectKey{
					// return the value without blank char
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(getValue("example.ini", "remote \"origin\"", "fetch"))
	fmt.Println(getValue("example.ini", "core", "hideDotFiles"))
}
