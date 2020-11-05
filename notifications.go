package main

// NotificationsService is an object that manages orders.
type NotificationsService struct {
	Orders            []Order
	stockCountApples  int
	stockCountOranges int
}

// GetOrder gets an order with the specified ID from the list of orders.
func (s *NotificationsService) GetOrder(id int) (order Order) {
	for _, o := range s.Orders {
		if o.ID == id {
			order = o
		}
	}
	return
}

// RegisterOrder adds a new order to the list of orders.
func (s *NotificationsService) RegisterOrder(order Order) (result Order) {
	s.Orders = append(s.Orders, order)
	result = s.GetOrder(order.ID)
	return
}
