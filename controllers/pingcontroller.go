package controllers

import "github.com/gin-gonic/gin"

// Ping godoc
// @Summary      Ping
// @Description  Ping system to check if it works
// @Tags         Diagnostic
// @Produce      json
// @Failure  	 401
// @Router       /ping [get]
// @Security ApiKeyAuth
func Ping(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}
