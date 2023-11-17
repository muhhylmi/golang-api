package web

import (
	"golang-api/modules/cart/models/domain"
	"golang-api/utils"

	"github.com/google/uuid"
)

type RequestCreateCart struct {
	Price float64 `json:"price"`

	Details []Details `json:"details"`
	Token   utils.ClaimToken
}

type Details struct {
	BookId string `json:"bookId"`
}

func (r *RequestCreateCart) ToModel() *domain.Cart {
	cartId := uuid.New().String()
	var detailCarts []domain.CartDetail
	for _, detail := range r.Details {
		detailCarts = append(detailCarts, domain.CartDetail{
			Id:        uuid.New().String(),
			CartId:    cartId,
			BookId:    detail.BookId,
			Status:    CREATED,
			CreatedBy: r.Token.UserId,
			UpdatedBy: r.Token.UserId,
		})
	}
	return &domain.Cart{
		Id:        cartId,
		UserId:    r.Token.UserId,
		CreatedBy: r.Token.UserId,
		UpdatedBy: r.Token.UserId,
		Details:   detailCarts,
	}
}
