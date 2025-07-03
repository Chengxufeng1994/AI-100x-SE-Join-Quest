package application

import (
	"github.com/Chengxufeng1994/ai-100x-se-join-quest/orderpromotions/internal/order/domain"
)

// BogoCosmeticsStrategy applies buy one get one for cosmetics.
type BogoCosmeticsStrategy struct{}

func (b *BogoCosmeticsStrategy) Apply(order *domain.Order) {
	for i, item := range order.Items {
		if item.Category == "cosmetics" {
			order.Items[i].Quantity = item.Quantity + 1
		}
	}
}

// Double11Strategy applies the Double 11 discount: every 10 items get 20% off.
type Double11Strategy struct{}

func (d *Double11Strategy) Apply(order *domain.Order) {
	total := 0
	for i, item := range order.Items {
		qty := item.Quantity
		set := qty / 10
		rest := qty % 10
		discounted := set * 10 * item.UnitPrice * 80 / 100
		undiscounted := rest * item.UnitPrice
		total += discounted + undiscounted
		order.Items[i].Quantity = qty // keep original quantity
	}
	order.TotalAmount = total
}

// ThresholdDiscountStrategy applies a discount if the original amount exceeds a threshold.
type ThresholdDiscountStrategy struct {
	Threshold int
	Discount  int
}

func (t *ThresholdDiscountStrategy) Apply(order *domain.Order) {
	if t.Threshold > 0 && t.Discount > 0 && order.OriginalAmount >= t.Threshold {
		order.Discount = t.Discount
		order.TotalAmount -= t.Discount
		if order.TotalAmount < 0 {
			order.TotalAmount = 0
		}
	}
}
