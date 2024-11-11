package product

import (
	"unicode/utf8"

	errorDomain "github.com/nara-ryoya/go-clean-architecture/app/domain/error"
	"github.com/nara-ryoya/go-clean-architecture/go-pkg/ulid"
)

type Product struct {
	id string
	ownerID string
	name string
	description string
	price int64
	stock int
}

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	if !uid.IsValid(ownerID) {
		return nil, errorDomain.NewError("オーナーIDの値が不正です")
	}
}