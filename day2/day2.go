package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var total int

func main() {
	file, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbewsString := strings.Fields(line)
		println(line)
		var oldnumbew *int
		safe := true
		for j := range numbewsString {

			intNumbew, err := strconv.Atoi(numbewsString[j])
			if err != nil {
				println("oh no")
			}
			//I dont really mind the overhead :3
			if oldnumbew == nil {
				oldnumbew = &intNumbew

			} else if -3 > (*oldnumbew-intNumbew) || (*oldnumbew-intNumbew) > 3 || (*oldnumbew-intNumbew) == 0 {
				safe = false
				print("this bad")
				break
			}
			oldnumbew = &intNumbew
		}
		if safe {
			total += 1
		}
	}
	print(total)
}
