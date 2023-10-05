package main

import (
	"log"
	"time"

	"ton-ton/database"
	"ton-ton/delivery/http"
	"ton-ton/delivery/http/middleware"
	"ton-ton/repository"
	"ton-ton/usecase"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("app.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := echo.New()
	router.Use(middleware.InitMiddleware().CORS)

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
	defer db.Close()

	repo := repository.New(db)
	uc := usecase.NewUsecase(repo, time.Second*2)

	group := router.Group("v1")
	http.NewHandler(group, uc)

	log.Fatal(router.Start(viper.GetString("address")))
}
