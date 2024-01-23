package web

import (
	"golang-api/modules/cart/models/domain"
)

type ResponseCart struct {
	UserId    string    `json:"userId"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt int64     `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	Details   []Details `json:"cartDetails"`
}

func ToResponseCart(carts []*domain.Cart) []ResponseCart {
	var totalPrice float64
	var response []ResponseCart
	for _, cart := range carts {
		var detailCarts []Details
		for _, detail := range cart.Details {
			totalPrice += (detail.Books.Price * float64(detail.Qty))
			detailCarts = append(detailCarts, Details{
				BookId:   detail.BookId,
				BookName: detail.Books.Title,
				Price:    detail.Books.Price,
				Qty:      detail.Qty,
			})
		}
		response = append(response, ResponseCart{
			UserId:    cart.UserId,
			Price:     totalPrice,
			Status:    cart.Status,
			CreatedAt: cart.CreatedAt,
			CreatedBy: cart.CreatedBy,
			Details:   detailCarts,
		})
	}
	return response
}
