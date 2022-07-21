package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	// open file
	file,err := os.Open("./input")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	task1(file)

	file2,err := os.Open("./input")
	if err != nil{
		panic(err)
	}
	defer file2.Close()
	task2(file2)

}

func task1(file *os.File){
	scanner := bufio.NewScanner(file)
	var x, y int
	for scanner.Scan(){
		split := strings.Split(scanner.Text(), " ")
		position,_ := strconv.Atoi(split[1])
		switch split[0]{
		case "up": y -= position
		case "down": y += position
		case "forward": x += position;
		}
		
	}
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(x*y)

}

func task2(file *os.File){
	var x, y, depth int
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		split := strings.Split(scanner.Text(), " ")
		position,_ := strconv.Atoi(split[1])
		switch split[0]{
		case "up": y -= position
		case "down": y += position
		case "forward": x += position; depth = y*position
		}
	}
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
	fmt.Println(x*depth)

}