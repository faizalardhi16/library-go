package main

import (
	"library/auth"
	"library/book"
	"library/handler"
	"library/helper"
	"library/models"
	"library/reservation"
	"library/transaction"
	"library/users"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/joybox?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	models.InitializedDB(db)

	//repository
	userRepository := users.NewUserRepository(db)
	NewBookRepository := book.NewBookRepository(db)
	transactionRepository := transaction.NewTransactionRepository(db)
	reservationRepository := reservation.NewReservationRepository(db)

	//service
	userService := users.NewUserService(userRepository)
	authService := auth.NewService()
	bookService := book.NewBookService(NewBookRepository)
	transactionService := transaction.NewTransactionService(transactionRepository)
	reservationService := reservation.NewReservationService(reservationRepository)

	//handler
	userHandler := handler.NewUserHandler(userService, authService)
	bookHandler := handler.NewBookHandler(bookService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	reservationHandler := handler.NewReservationHandler(reservationService, transactionService)

	router := gin.Default()

	api := router.Group("/api/v1")

	//get request
	api.GET("/books", bookHandler.FindBook)
	api.GET("/transaction", transactionHandler.GetTransaction)

	//post request
	api.POST("/register", userHandler.RegisterUserHandler)
	api.POST("/check-email", userHandler.CheckEmailAvailability)
	api.POST("/session", userHandler.LoginUser)
	api.POST("/books", bookHandler.CreateBook)
	api.POST("/books/:id", bookHandler.UploadFile)
	api.POST("/transaction", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/reservation", authMiddleware(authService, userService), reservationHandler.CreateReservation)

	//put request

	router.Run(":5041")

}

func authMiddleware(authService auth.Service, userService users.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("An unauthorized 1", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""

		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("An unauthorized 2", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("An unauthorized 3", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("An unauthorized 4", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("CurrentUser", user)
	}

}
