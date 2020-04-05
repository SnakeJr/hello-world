package main

import "fmt"

func main() {
	printSomething("anything")
	printSomething("somethingelse")
	printSomething("morewords")
}

func printSomething(natalie string) {
	fmt.Println(natalie)
	var input string
	fmt.Scanln(&input)
	natalie += input
	fmt.Println(fmt.Sprintf("%s is what you told me, and i was called with %s", input, natalie))
}
