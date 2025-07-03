package test

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/ai-100x-se-join-quest/orderpromotions/internal/order/application"
	"github.com/Chengxufeng1994/ai-100x-se-join-quest/orderpromotions/internal/order/domain"
	"github.com/cucumber/godog"
)

var (
	orderService   *application.OrderService
	orderItems     []domain.OrderItem
	orderResult    *domain.Order
	threshold      int
	discount       int
	bogoActive     bool
	double11Active bool
)

func resetScenarioVars() {
	orderItems = nil
	orderResult = nil
	threshold = 0
	discount = 0
	bogoActive = false
	double11Active = false
}

func buildOrderService() *application.OrderService {
	var strategies []application.DiscountStrategy
	if bogoActive {
		strategies = append(strategies, &application.BogoCosmeticsStrategy{})
	}
	if double11Active {
		strategies = append(strategies, &application.Double11Strategy{})
	}
	if threshold > 0 && discount > 0 {
		strategies = append(strategies, &application.ThresholdDiscountStrategy{Threshold: threshold, Discount: discount})
	}
	return &application.OrderService{Strategies: strategies}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		resetScenarioVars()
		return ctx, nil
	})
	ctx.Step(`^the double eleven promotion is active$`, func() error {
		double11Active = true
		return nil
	})
	ctx.Step(`^no promotions are applied$`, func() error {
		bogoActive = false
		double11Active = false
		threshold = 0
		discount = 0
		return nil
	})
	ctx.Step(`^the threshold discount promotion is configured:$`, func(table *godog.Table) error {
		for i, row := range table.Rows {
			if i == 0 { // skip header
				continue
			}
			threshold = atoi(row.Cells[0].Value)
			discount = atoi(row.Cells[1].Value)
		}
		return nil
	})
	ctx.Step(`^the buy one get one promotion for cosmetics is active$`, func() error {
		bogoActive = true
		return nil
	})
	ctx.Step(`^a customer places an order with:$`, func(table *godog.Table) error {
		orderItems = nil
		headers := table.Rows[0].Cells
		for i, row := range table.Rows {
			if i == 0 { // skip header
				continue
			}
			item := domain.OrderItem{}
			for j, h := range headers {
				switch h.Value {
				case "productName":
					item.ProductName = row.Cells[j].Value
				case "category":
					item.Category = row.Cells[j].Value
				case "quantity":
					item.Quantity = atoi(row.Cells[j].Value)
				case "unitPrice":
					item.UnitPrice = atoi(row.Cells[j].Value)
				}
			}
			orderItems = append(orderItems, item)
		}
		orderService = buildOrderService()
		orderResult = orderService.Checkout(orderItems)
		return nil
	})
	ctx.Step(`^the order summary should be:$`, func(table *godog.Table) error {
		if orderResult == nil {
			return fmt.Errorf("orderResult is nil, likely due to missing or failed When step")
		}
		headers := table.Rows[0].Cells
		values := table.Rows[1].Cells
		for i, h := range headers {
			expected := atoi(values[i].Value)
			switch h.Value {
			case "originalAmount":
				if orderResult.OriginalAmount != expected {
					return fmt.Errorf("originalAmount not match: expected %d, got %d", expected, orderResult.OriginalAmount)
				}
			case "discount":
				if orderResult.Discount != expected {
					return fmt.Errorf("discount not match: expected %d, got %d", expected, orderResult.Discount)
				}
			case "totalAmount":
				if orderResult.TotalAmount != expected {
					return fmt.Errorf("totalAmount not match: expected %d, got %d", expected, orderResult.TotalAmount)
				}
			}
		}
		return nil
	})
	ctx.Step(`^the customer should receive:$`, func(table *godog.Table) error {
		if orderResult == nil {
			return fmt.Errorf("orderResult is nil, likely due to missing or failed When step")
		}
		if len(orderResult.Items) != len(table.Rows)-1 {
			return fmt.Errorf("expected %d items, got %d", len(table.Rows)-1, len(orderResult.Items))
		}
		for i, row := range table.Rows {
			if i == 0 { // skip header
				continue
			}
			item := orderResult.Items[i-1]
			if item.ProductName != row.Cells[0].Value {
				return fmt.Errorf("expected product %s, got %s", row.Cells[0].Value, item.ProductName)
			}
			expectedQty := atoi(row.Cells[1].Value)
			if item.Quantity != expectedQty {
				return fmt.Errorf("expected quantity %d, got %d", expectedQty, item.Quantity)
			}
		}
		return nil
	})
}

func atoi(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
