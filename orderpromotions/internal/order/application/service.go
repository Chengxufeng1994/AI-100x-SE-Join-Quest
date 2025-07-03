package application

import (
	"github.com/Chengxufeng1994/ai-100x-se-join-quest/orderpromotions/internal/order/domain"
)

// DiscountStrategy defines the interface for all discount/price rules.
type DiscountStrategy interface {
	Apply(order *domain.Order) // modifies order in place
}

// OrderService handles order creation and discount logic.
type OrderService struct {
	Strategies []DiscountStrategy
}

// Checkout processes the order items and returns an Order.
func (s *OrderService) Checkout(items []domain.OrderItem) *domain.Order {
	// Calculate original amount
	original := 0
	for _, item := range items {
		original += item.UnitPrice * item.Quantity
	}
	order := &domain.Order{
		Items:          make([]domain.OrderItem, len(items)),
		OriginalAmount: original,
		Discount:       0,
		TotalAmount:    original,
	}
	copy(order.Items, items)

	// Apply all discount strategies in order
	for _, strategy := range s.Strategies {
		strategy.Apply(order)
	}

	return order
}
