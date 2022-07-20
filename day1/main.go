package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	// open file
	file,err := os.Open("./input")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	// scan file into slice of ints
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	var increment int
	// start from the 1 item
	for i := range numbers[1:]{
		if numbers[i] < numbers[i+1]{
			increment++
		}
	}
	fmt.Println("How many measurements are larger than the previous measurement?")
	fmt.Println(increment)
}