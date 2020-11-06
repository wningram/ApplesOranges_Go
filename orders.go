package main

import (
	"fmt"
	"strings"
)

var totalOrders int = 0

// Order holds information about orders placed.
type Order struct {
	ID         int
	Status     bool
	OrderInput []string
}

// GetNewID assigns an ID to the current Order by incrementing the total
// number of orders created.
func (o *Order) GetNewID() int {
	totalOrders++
	o.ID = totalOrders
	return o.ID
}

// CalculateDiscount determines the discount to apply to teh overall order. Value is negative.
func (o *Order) CalculateDiscount() float64 {
	var totalDiscount = 0.0
	var appleCount = 0
	var orangeCount = 0

	for _, itm := range o.OrderInput {
		switch strings.ToLower(itm) {
		case "apple":
			appleCount++
		case "orange":
			orangeCount++
		}
	}

	totalDiscount += float64(appleCount/2) * 0.60
	totalDiscount += float64(orangeCount/3) * 0.25

	return totalDiscount * -1
}

// CalculateGrossCost calculates teh cost of a list of products.
func (o *Order) CalculateGrossCost() (totalCost float64) {
	for _, itm := range o.OrderInput {
		switch strings.ToLower(itm) {
		case "apple":
			totalCost += 0.60
		case "orange":
			totalCost += 0.25
		}
	}
	return
}

func (o Order) String() string {
	return fmt.Sprintf("{\nOrder Id: %d\nStatus: %t\nInputs: %v\n}",
		o.ID, o.Status, o.OrderInput)
}
