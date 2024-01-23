package web

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/utils/jwt"

	"github.com/google/uuid"
)

type RequestCreateCart struct {
	Price float64 `json:"price"`

	Details []Details `json:"details" validate:"required"`
	Token   jwt.ClaimToken
}

type RequestListCart struct {
	Status string `params:"status" query:"status" validate:"required,oneof=created paid"`
}

type Details struct {
	BookId   string  `json:"bookId" validate:"required"`
	Price    float64 `json:"price"`
	BookName string  `json:"bookName"`
	Qty      int     `json:"qty"`
}

func (r *RequestCreateCart) ToModel() *domain.Cart {
	cartId := uuid.New().String()
	var detailCarts []domain.CartDetail
	for _, detail := range r.Details {
		detailCarts = append(detailCarts, domain.CartDetail{
			Id:        uuid.New().String(),
			CartId:    cartId,
			BookId:    detail.BookId,
			Qty:       detail.Qty,
			CreatedBy: r.Token.UserId,
			UpdatedBy: r.Token.UserId,
		})
	}
	return &domain.Cart{
		Id:        cartId,
		UserId:    r.Token.UserId,
		Status:    CREATED,
		CreatedBy: r.Token.UserId,
		UpdatedBy: r.Token.UserId,
		Details:   detailCarts,
	}
}
