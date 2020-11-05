package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	helpText := "Enter a comma separated list of items. Items can be: Apple or Orange." +
		"\nTo restock enter restock apples/oranges ENTER. THen type the number of stock to add."
	run := true
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("App is running!")

	for run {
		fmt.Println("Input -> ")
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("An error has occured: %v", err.Error())
			break
		}

		switch strings.ToLower(string(input)) {
		case "help":
			fmt.Println(helpText)
		case "quit":
			run = false
		default:
			break
		}
	}

	fmt.Println("Program sucessfully terminated.")
}
