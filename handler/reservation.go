package handler

import (
	"library/helper"
	"library/reservation"
	"library/transaction"
	"library/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reservationHandler struct {
	reservationService reservation.Service
	transactionService transaction.Service
}

func NewReservationHandler(reservationService reservation.Service, transactionService transaction.Service) *reservationHandler {
	return &reservationHandler{reservationService, transactionService}
}

func (h *reservationHandler) CreateReservation(c *gin.Context) {
	var input reservation.ReservationInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Create reservation failed! 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.BookingDate = "30-11-2022"
	currentUser := c.MustGet("CurrentUser").(users.User)
	input.UserID = currentUser.ID

	newRes, err := h.reservationService.CreateReservation(input)

	if err != nil {
		response := helper.APIResponse("Create reservation failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, id := range input.TransactionID {
		inputData := transaction.UpdateTransactionInput{}
		inputData.ID = id
		inputData.ReservationID = newRes.ID
		_, err := h.transactionService.UpdateTransaction(inputData)

		if err != nil {
			response := helper.APIResponse("Create reservation & transaction failed!", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := reservation.ResponseReservationFormat(newRes)

	response := helper.APIResponse("Create Reservation Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
