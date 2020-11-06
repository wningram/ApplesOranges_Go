package main

import (
	"fmt"
	"strings"
)

// NotificationsService is an object that manages orders.
type NotificationsService struct {
	Orders            []Order
	stockCountApples  int
	stockCountOranges int
}

func (s *NotificationsService) ReportStockCounts() {
	fmt.Println("Current Stock Counts:")
	fmt.Printf("Total Apples: %d\n", s.stockCountApples)
	fmt.Printf("Total Oranges: %d\n\n", s.stockCountOranges)
}

// GetOrder gets an order with the specified ID from the list of orders.
func (s *NotificationsService) GetOrder(id int) (order *Order, err error) {
	for i := 0; i < len(s.Orders); i++ {
		if s.Orders[i].ID == id {
			order = &s.Orders[i]
			err = nil
			return
		}
	}

	err = fmt.Errorf("Item with ID %d does not exist", id)
	return
}

// RegisterOrder adds a new order to the list of orders.
func (s *NotificationsService) RegisterOrder(order Order) (result *Order) {
	order.GetNewID()
	s.Orders = append(s.Orders, order)
	result, _ = s.GetOrder(order.ID)
	return
}

// ProcessOrder runs processing logic on teh order with teh specified ID.
func (s *NotificationsService) ProcessOrder(id int) error {
	order, err := s.GetOrder(id)
	applesAvail := s.stockCountApples
	orangesAvail := s.stockCountOranges
	if err != nil {
		return err
	}

	// Determine if order can be fulfilled with current stock
	for _, itm := range order.OrderInput {
		switch strings.ToLower(itm) {
		case "apple":
			if applesAvail--; applesAvail == -1 {
				return ErrOutOfStock{"Apple"}
			}
		case "orange":
			if orangesAvail--; orangesAvail == -1 {
				return ErrOutOfStock{"Orange"}
			}
		}
	}

	// Update stock count for this order
	s.stockCountApples -= s.stockCountApples - applesAvail
	s.stockCountOranges -= s.stockCountOranges - orangesAvail
	// Indicate that this order has been processed
	order.Status = true
	return nil
}

// GetDeliveryTime gets the estimated delivery time for an order based on its
// position in teh orders list and how many orders have already been processed.
func (s *NotificationsService) GetDeliveryTime(orderID int) (deliveryTime int, err error) {
	currentOrder, err := s.GetOrder(orderID)
	if err != nil {
		return -1, err
	}

	if currentOrder.Status {
		return 0, nil
	}
	for _, o := range s.Orders {
		if !o.Status && o.ID < orderID {
			deliveryTime += 10 * 60
		}
	}
	return
}

// ReprocessUnfulfilledOrders runs processing logic on all orders that have not yet been
// fulfilled.
func (s *NotificationsService) ReprocessUnfulfilledOrders() (processedOrders []Order) {
	for i := 0; i < len(s.Orders); i++ {
		if !s.Orders[i].Status {
			fmt.Printf("Current Order: %v", s.Orders[i])
			fmt.Printf("Processing order ID %d\n", s.Orders[i].ID) // XXX
			s.ProcessOrder(s.Orders[i].ID)
			processedOrders = append(processedOrders, s.Orders[i])
		}
	}
	return
}

// RestockApples increments apple stock by specified amount and attempts
// to reprocess unfulfilled orders. Returns a list of fulfilled orders.
func (s *NotificationsService) RestockApples(count int) []Order {
	s.stockCountApples += count
	return s.ReprocessUnfulfilledOrders()
}

// RestockOranges increments orange stock by specified amount and attempts
// to reprocess unfulfilled orders. Returns fulfilled orders.
func (s *NotificationsService) RestockOranges(count int) []Order {
	s.stockCountOranges += count
	return s.ReprocessUnfulfilledOrders()
}
