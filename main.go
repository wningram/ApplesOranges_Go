package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateInput(inputs []string) error {
	for _, input := range inputs {
		switch strings.ToLower(input) {
		case "apple":
		case "orange":
		default:
			return fmt.Errorf("Invalid input")
		}
	}
	return nil
}

func parseInput(input string) ([]string, error) {
	data := strings.ReplaceAll(input, " ", "")
	words := strings.Split(data, ",")
	err := validateInput(words)
	return words, err
}

func main() {
	helpText := "Enter a comma separated list of items. Items can be: Apple or Orange." +
		"\nTo restock enter restock apples/oranges ENTER. THen type the number of stock to add."
	run := true
	reader := bufio.NewReader(os.Stdin)
	notifications := NotificationsService{stockCountApples: 3, stockCountOranges: 3}
	fmt.Println("App is running!")

	for run {
		fmt.Println("Input -> ")
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("An error has occured: %v\n", err.Error())
			break
		}

		// fmt.Printf("Input was: %v %T\n", strings.ToLower(string(input)), string(input))
		switch strings.ToLower(string(input)) {
		case "help":
			fmt.Println(helpText)
		case "quit":
			run = false
		case "orders":
			for _, o := range notifications.Orders {
				fmt.Println(o.String())
			}
		case "restock apples":
			count, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("There was an error parsing input: %v\n", err.Error())
			}
			// Convert input to int
			restockCount, err := strconv.Atoi(count)
			if err != nil {
				fmt.Printf("The value provided is not a valid number.\n")
			}
			processedOrders := notifications.RestockApples(restockCount)
			for _, o := range processedOrders {
				fmt.Printf("Order %d processed.\n", o.ID)
				fmt.Println(o)
			}
		case "restock oranges":
			count, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("There was an error parsing input: %v\n", err.Error())
			}
			// Convert input to int
			restockCount, err := strconv.Atoi(count)
			if err != nil {
				fmt.Println("The value provided is not a valid number.")
			}
			processedOrders := notifications.RestockOranges(restockCount)
			for _, o := range processedOrders {
				fmt.Printf("Order %d processed.\n", o.ID)
				fmt.Println(o)
			}
		default:
			orderInput, err := parseInput(string(input))
			if err != nil {
				fmt.Println(err.Error())
			}
			newOrder := notifications.RegisterOrder(Order{OrderInput: orderInput})
			discount := newOrder.CalculateDiscount()
			cost := newOrder.CalculateGrossCost() + discount
			err = notifications.ProcessOrder(newOrder.ID)
			oosErr, ok := err.(ErrOutOfStock)
			if ok {
				deliveryTime, err := notifications.GetDeliveryTime(newOrder.ID)
				if err != nil {
					fmt.Println("Order does not exist.")
					break
				}
				fmt.Printf("Could not process order, the following is out of stock: %v\n", oosErr.product)
				fmt.Printf("Estimated Delivery Time: %d\n", deliveryTime)
			} else {
				deliveryTime, err := notifications.GetDeliveryTime(newOrder.ID)
				if err != nil {
					fmt.Println("Order does not exist.")
					break
				}
				fmt.Printf("Order %d processed.\n", newOrder.ID)
				fmt.Printf("Discount %f\n", discount)
				fmt.Printf("Total Cost is %f\n", cost)
				fmt.Printf("Estimated Delivery Time %ds (when in stock)\n", deliveryTime)
			}
		}
	}

	fmt.Println("Program sucessfully terminated.")
}
