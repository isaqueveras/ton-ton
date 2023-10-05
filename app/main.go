package main

import (
	"log"
	"time"

	"ton-ton/database"
	"ton-ton/delivery/http"
	"ton-ton/repository"
	"ton-ton/usecase"
	"ton-ton/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"
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

	secret := utils.Pointer(viper.GetString("secret"))
	cors := cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     viper.GetStringSlice("allow_origins"),
		AllowOriginFunc:  func(string) bool { panic("not implemented") },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
		AllowWildcard:    true,
	})

	router := gin.New()
	router.Use(cors, gopowersso.Authorization(secret), gopowersso.OnlyAdmin())

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
