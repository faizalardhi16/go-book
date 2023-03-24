package main

import (
	"go-book/auth"
	"go-book/author"
	"go-book/book"
	"go-book/category"
	"go-book/handler"
	"go-book/helper"
	"go-book/models"
	"go-book/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnect := envs["DB_SOURCE"]

	dsn := dbConnect

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot Connect to DB")
	}

	models.RegistryDatabase(db)

	//repository
	userRepo := user.NewRepositoryUser(db)
	categoryRepo := category.NewCategoryRepository(db)
	authorRepo := author.NewAuthorRepository(db)
	bookRepo := book.NewBookRepository(db)

	//service
	userService := user.NewServiceUser(userRepo)
	authService := auth.NewService()
	categoryService := category.NewCategoryService(categoryRepo)
	authorService := author.NewAuthorService(authorRepo)
	bookService := book.NewBookService(bookRepo, authorRepo, categoryRepo)

	//handler
	userHandler := handler.NewHandlerUser(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	authorHandler := handler.NewAuthorHandler(authorService)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Access-Control-Allow-Headers", "Accept", "XMLHttpRequest"},
		ExposeHeaders:   []string{"Data-Length"},
		AllowAllOrigins: true,
		MaxAge:          12 * time.Hour,
	}))

	api := router.Group("/api/v1")

	//GET
	api.GET(("/"), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "HELLO")
	})

	//POST
	api.POST("/user", userHandler.RegisterUserHandler)
	api.POST("/email-check", userHandler.CheckEmailAvailibility)
	api.POST("/login", userHandler.LoginUserHandler)
	api.POST("/category", authMiddleware(authService, userService), categoryHandler.CreateCategoryHandler)
	api.POST("/author", authMiddleware(authService, userService), authorHandler.CreateAuthorHandler)
	api.POST("/book", authMiddleware(authService, userService), bookHandler.CreateBookHandler)

	//DELETE
	api.DELETE("/author", authMiddleware(authService, userService), authorHandler.DeleteAuthorHandler)

	//GET
	api.GET("/book", bookHandler.GetAllBook)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
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

		user := claim["user"].(interface{})

		c.Set("CurrentUser", user)

	}
}
