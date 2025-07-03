package domain

// Order represents a customer order.
type Order struct {
	Items          []OrderItem
	OriginalAmount int
	Discount       int
	TotalAmount    int
}

// OrderItem represents an item in an order.
type OrderItem struct {
	ProductName string
	Category    string
	Quantity    int
	UnitPrice   int
}

// Product represents a product for sale.
type Product struct {
	Name      string
	Category  string
	UnitPrice int
}
