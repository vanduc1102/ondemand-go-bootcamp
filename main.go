package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/controllers"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/utils"
)

func main() {

	_, err := utils.LoadConfig(".", "local")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"at":     time.Now(),
		})
	})

	router.GET("pokemon/", controllers.FindPokemonList)
	router.GET("pokemon/:id", controllers.FindPokemon)
	router.Run("localhost:8080")
}
