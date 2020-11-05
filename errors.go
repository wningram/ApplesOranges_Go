package main

// ErrOutOfStock is thrown when a product is out of stock and an order
// cannot be fulfilled.
type ErrOutOfStock struct {
	product string
}

func (e ErrOutOfStock) Error() string {
	return e.product
}
