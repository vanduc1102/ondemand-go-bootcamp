package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/services"
)

func FindPokemonList(ctx *gin.Context) {
	pokemonList, error := services.FindAll()
	if error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": pokemonList})
}

func FindPokemon(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid pokemon id=%s", paramId)})
		return
	}

	pokemon, error := services.FindById(id)
	if error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": pokemon,
	})
}

type ImportPokemonInput struct {
	Limit  int `json:"limit" binding:"required,min=1"`
	Offset int `json:"offset" binding:"required,min=0"`
}

func Import(ctx *gin.Context) {
	var input ImportPokemonInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body, error := services.Import(input.Limit, input.Offset)
	if error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": body,
	})
}
