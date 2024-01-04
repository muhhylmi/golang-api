package handler

// import (
// 	"context"
// 	"fmt"
// 	"golang-api/modules/books/models/web"
// 	models "golang-api/modules/books/models/web"
// 	"golang-api/proto"
// 	"golang-api/utils"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func (h *HTTPHandler) GrpcCreateBook(c context.Context, req *proto.BookDataRequest) *proto.BookDataResponse {
// 	payload := web.RequestCreateBook{
// 		Title:  req.Name,
// 		Author: req.Author,
// 	}

// 	fmt.Println(payload)
// 	return &proto.BookDataResponse{Success: true}
// }

// func (h *HTTPHandler) CreateBookByGrpc(c echo.Context) error {
// 	log := utils.LogWithContext(h.logger, contextName, "CreateBook")
// 	book := new(models.RequestCreateBook)
// 	book.Token = c.Get("user").(utils.ClaimToken)
// 	if err := utils.BindValidate(c, book); err != nil {
// 		log.Error(err)
// 		perr := utils.ResultFailed(utils.NewBadRequest(err.Error()), utils.ValidationError)
// 		return utils.ResponseError(perr.Error, perr.StatusCode, c)
// 	}

// 	result := h.usecase.CreateBook(c.Request().Context(), book)
// 	if result.Error != nil {
// 		return utils.ResponseError(result.Error, result.StatusCode, c)
// 	}

// 	return utils.Response(result.Data, "Your Request has been Approve", http.StatusCreated, c)
// }
