//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
	"time"

	cartDomain "github.com/nara-ryoya/go-clean-architecture/app/domain/cart"
	errorDomain "github.com/nara-ryoya/go-clean-architecture/app/domain/error"
	"github.com/nara-ryoya/go-clean-architecture/pkg/ulid"
)

type Order struct {
	id string
	userID string
	totalAmount int64
	products OrderProducts
	orderedAt time.Time
}

func (p *Order) ID() string {
	return p.id
}

func (p *Order) UserID() string {
	return p.userID
}

func (p *Order) TotalAmount() int64 {
	return p.totalAmount
}

func (p *Order) Products() []OrderProduct {
	return p.products
}

func (p *Order) OrderedAt() time.Time {
	return p.orderedAt
}

func (p *Order) ProductIDs() []string {
	var productIDs []string
	for _, product := range p.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

type OrderProducts []OrderProduct

type OrderProduct struct {
	productID string
	price     int64
	quantity  int
}

func newOrder(
	id string,
	userID string,
	totalAmount int64,
	products []OrderProduct,
	orderedAt time.Time,
) (*Order, error) {
	if !ulid.IsValid(userID) {
		return nil, errorDomain.NewError("ユーザーIDの値が不正です。")
	}
	if totalAmount < 0 {
		return nil, errorDomain.NewError("合計金額の値が不正です。")
	}
	if len(products) < 1 {
		return nil, errorDomain.NewError("商品がありません。")
	}

	return &Order{
		id: id,
		userID: userID,
		totalAmount: totalAmount,
		products: products,
		orderedAt: orderedAt,
	}, nil
}

func NewOrder(userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	return newOrder(
		ulid.NewULID(),
		userID,
		totalAmount,
		products,
		now,
	)
}

func Reconstruct(id string, userID string, totalAmount int64, products []OrderProduct, OrderedAt time.Time) (*Order, error) {
	return newOrder(
		id,
		userID,
		totalAmount,
		products,
		OrderedAt,
	)
}

func (p OrderProducts) ProductIDs() []string {
	var productIDs []string
	for _, product := range p {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

func (p OrderProducts) TotalAmount() int64 {
	var totalAmount int64
	for _, product := range p {
		totalAmount += product.price * int64(product.quantity)
	}
	return totalAmount
}

func (p *OrderProduct) ProductID() string {
	return p.productID
}

func (p *OrderProduct) Quantity() int {
	return p.quantity
}

func (p *OrderProduct) Price() int64 {
	return p.price
}

type OrderDomainService interface {
	OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error)
}
