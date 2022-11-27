package handler

import (
	"fmt"
	"library/book"
	"library/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var books book.CreateBookInput

	err := c.ShouldBindJSON(&books)

	if err != nil {
		response := helper.APIResponse("Create book failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newBook, err := h.bookService.CreateBook(books)

	if err != nil {
		response := helper.APIResponse("Create book failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := book.CreateBookFormatter(newBook)

	response := helper.APIResponse("Created book was sucessfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) FindBook(c *gin.Context) {

	books, err := h.bookService.GetBook()

	if err != nil {
		response := helper.APIResponse("Create book failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Load book was sucessfully", http.StatusOK, "success", book.BookFormat(books))

	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) UploadFile(c *gin.Context) {
	var input book.UpdateInput
	file, err := c.FormFile("file_name")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to Upload file Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindUri(&input)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to Upload file Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%s-%s", uuid.New().String(), file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to Upload file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.bookService.UpdateBook(input.ID, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to Upload file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.bookService.GetBookByID(input.ID)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to Upload file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}

	response := helper.APIResponse("Success to upload file", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
