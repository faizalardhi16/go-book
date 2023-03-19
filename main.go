package main

import (
	"go-book/handler"
	"go-book/models"
	"go-book/user"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnect := envs["GO_CONNECT_DB"]

	dsn := dbConnect

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot Connect to DB")
	}

	models.RegistryDatabase(db)

	//repository
	userRepo := user.NewRepositoryUser(db)

	//service
	userService := user.NewServiceUser(userRepo)

	//handler
	userHandler := handler.NewHandlerUser(userService)

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

	router.Run()

}
