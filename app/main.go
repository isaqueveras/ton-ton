package main

import (
	"log"
	"time"

	"ton-ton/database"
	"ton-ton/delivery/http"
	"ton-ton/repository"
	"ton-ton/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           []string{},
		AllowOriginFunc:        func(string) bool { panic("not implemented") },
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:           []string{},
		AllowCredentials:       false,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
	defer db.Close()

	repo := repository.New(db)
	uc := usecase.NewUsecase(repo, time.Second*2)

	group := router.Group("v1")
	http.NewHandler(group, uc)

	log.Fatal(router.Run(viper.GetString("address")))
}
