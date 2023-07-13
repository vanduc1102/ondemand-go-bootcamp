package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/services"
)

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
