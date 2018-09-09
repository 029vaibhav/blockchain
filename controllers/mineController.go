package controllers

import (
	"bitbucket.org/blockchain/services"
	"github.com/labstack/echo"
	"net/http"
)

func (c *Controller) MineTransactions(ctx echo.Context) error {
	block := services.MineTransaction()
	return ctx.JSON(http.StatusOK, GetSuccessResponse(block))
}
