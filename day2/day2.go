package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var total int
var safe bool = true
var oldnumbew *int

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var intList []int
		line := scanner.Text()
		numbewsString := strings.Fields(line)

		safe = true

		//conv list string to list int
		for j := range numbewsString {
			intNumbew, err := strconv.Atoi(numbewsString[j])
			if err != nil {
				println("oh no")
			}
			intList = append(intList, intNumbew)
		}

		part2(intList)
	}

	print(total)
}

func part1(intList []int) {
	oldnumbew = &intList[0]
	isInc := false
	if intList[0] < intList[1] {
		isInc = true
	} else if intList[0] == intList[1] {
		return
	}

	for j := range intList {
		newnumbew := intList[j]
		if j > 0 {
			if isInc {
				if -3 > (*oldnumbew-newnumbew) || newnumbew <= *oldnumbew {
					safe = false
					break
				}
			} else {
				if (*oldnumbew-newnumbew) > 3 || newnumbew >= *oldnumbew {
					safe = false
					break
				}
			}
		}
		oldnumbew = &newnumbew
	}
	if safe {
		total += 1
	}
}

func part2(intList []int) {
	oldnumbew = &intList[0]
	safeOnce := true
	isInc := false
	startfrom := 0
	if intList[0] < intList[1] {
		isInc = true
	} else if intList[0] == intList[1] {
		safeOnce = false
		startfrom = 1
		intList[1] = intList[0]
		if intList[0] < intList[2] {
			isInc = true
		} else if intList[0] == intList[2] {
			return
		}
	}

	for j := range intList {
		newnumbew := intList[j]
		if j > startfrom {
			if isInc {
				if -3 > (*oldnumbew-newnumbew) || newnumbew <= *oldnumbew {
					if !safeOnce {
						safe = false
						break
					} else {
						safeOnce = false
						continue
					}
				}
			} else {
				if (*oldnumbew-newnumbew) > 3 || newnumbew >= *oldnumbew {
					if !safeOnce {
						safe = false
						break
					} else {
						safeOnce = false
						continue
					}
				}
			}
		}
		oldnumbew = &newnumbew
	}
	if safe {
		total += 1
	}
}
