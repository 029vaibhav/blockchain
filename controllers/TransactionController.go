package controllers

import (
	"bitbucket.org/blockchain/dto"
	. "bitbucket.org/blockchain/services"
	"bitbucket.org/blockchain/transaction"
	"github.com/labstack/echo"
	"net/http"
)

func (c *Controller) CreateTransaction(ctx echo.Context) error {

	req := dto.TransactionReq{}
	ctx.Bind(&req)
	createTransaction, e := CreateTransaction(req)
	if e == nil {
		return ctx.JSON(http.StatusOK, GetSuccessResponse(createTransaction))
	} else {
		return ctx.JSON(http.StatusBadRequest, GetBadResponse(e.Error()))

	}

}

func (c *Controller) GetWallet(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, GetSuccessResponse(GetWallet()))
}

func (c *Controller) GetTransactions(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, GetSuccessResponse(transaction.GetTransactions()))
}
