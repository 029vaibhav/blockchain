package controllers

import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/services"
	"github.com/labstack/echo"
	"net/http"
)

func (c *Controller) GetBlock(ctx echo.Context) error {
	chain := services.GetBlockChain()
	return ctx.JSON(http.StatusOK, GetSuccessResponse(chain))
}

func (c *Controller) AddBlock(ctx echo.Context) error {

	newBlock := block.Block{}
	ctx.Bind(&newBlock)
	if newBlock.Data == "" {
		return ctx.JSON(http.StatusBadRequest, GetBadResponse("data can not be empty"))
	}
	services.AddBlockChain(newBlock)
	return ctx.JSON(http.StatusOK, GetSuccessResponse(services.GetBlockChain()))
}
