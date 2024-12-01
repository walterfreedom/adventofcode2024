package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int

func main() {
	// URL to fetch the data
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//we need scanner ^^
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		// we wanna SPLIT line, we dont want "1905  1961"
		//apperantly this does exactly what we need !!
		numbers := strings.Fields(line)

		//wololo the string into int
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 == nil && err2 == nil {

			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
		//I do not care about handling the errors. this is a statement
	}

	part1()
	part2()

}

func part1() {
	//now at this step we should have two unsorted list.
	//is easy to sort tho :3
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] > list1[j]
	})

	//same for list 2
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] > list2[j]
	})

	//now they should be properly sorted 🗿🗿🗿
	//now I will just iterate and do the math :3 idk if there
	//are any better ways.
	distsum := 0
	for i := 0; i < len(list1); i++ {
		num1 := list1[i]
		num2 := list2[i]
		sum := num1 - num2
		if sum < 0 {
			sum *= -1
		}
		distsum += sum
	}
}

func part2() {
	similarity := 0
	//we use maps
	mp := make(map[int]int)
	for _, num := range list2 {
		mp[num] += 1
	}

	//maps are cool
	for _, num := range list1 {
		similarity += num * mp[num]
		fmt.Println(similarity)
	}
	//I like maps now
}
