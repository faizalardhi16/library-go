package handler

import (
	"library/helper"
	"library/transaction"
	"library/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(t transaction.Service) *transactionHandler {
	return &transactionHandler{t}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateInputValue
	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Create Transaction failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(users.User)

	inputValue := transaction.CreateTransactionInput{}
	inputValue.BookID = input.BookID
	inputValue.UserID = currentUser.ID

	trans, err := h.transactionService.CreateTransaction(inputValue)

	if err != nil {
		response := helper.APIResponse("Create Transaction failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.ResponseTransactionFormat(trans)

	response := helper.APIResponse("Create Transaction success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	trans, err := h.transactionService.GetTransaction()

	if err != nil {
		response := helper.APIResponse("Load Transaction failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.TransactionAllFormat(trans)

	response := helper.APIResponse("Load Transaction success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
