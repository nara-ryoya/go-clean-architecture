package product

import (
	"unicode/utf8"

	errorDomain "github.com/nara-ryoya/go-clean-architecture/app/domain/error"
	"github.com/nara-ryoya/go-clean-architecture/pkg/ulid"
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) OwnerID() string {
	return p.ownerID
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	if !ulid.IsValid(ownerID) {
		return nil, errorDomain.NewError("オーナーIDの値が不正です")
	}
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errorDomain.NewError("商品名の値が不正です")
	}
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errorDomain.NewError("商品説明の値が不正です")
	}
	if price < 1 {
		return nil, errorDomain.NewError("価格の値が不正です")
	}
	if stock < 0 {
		return nil, errorDomain.NewError("価格の値が不正です")
	}

	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func Reconstruct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(id, ownerID, name, description, price, stock)
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		ulid.NewULID(),
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

const (
	nameLengthMin        = 1
	nameLengthMax        = 100
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)
