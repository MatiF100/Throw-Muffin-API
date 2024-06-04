package controllers

import (
	"github.com/MatiF100/Throw-Muffin-API/database"
	"github.com/MatiF100/Throw-Muffin-API/models"
	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *RegisterUserRequest) toModel() models.User {
	return models.User{
		Email:    m.Email,
		Username: m.Username,
		Password: m.Password,
	}
}

// RegisterUser godoc
// @Summary      Register user
// @Tags         Auth
// @Param        request   body      RegisterUserRequest  true  "User data"
// @Produce      json
// @Failure  	 400
// @Failure  	 500
// @Success 	 200
// @Router       /auth/register [post]
func RegisterUser(context *gin.Context) {
	var userRequest RegisterUserRequest
	if err := context.ShouldBindJSON(&userRequest); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := userRequest.toModel()
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(400, gin.H{"error": "Error hashing password"})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
