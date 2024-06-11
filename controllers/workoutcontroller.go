package controllers

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/MatiF100/Throw-Muffin-API/database"
	"github.com/MatiF100/Throw-Muffin-API/models"
	tokenService "github.com/MatiF100/Throw-Muffin-API/services"
	"github.com/gin-gonic/gin"
)

type WorkoutPlanDetails struct {
	Excercises  []*models.Excercise
	DateCreated time.Time
}

// GenerateWorkout godoc
// @Summary      Generate workout plan
// @Tags         Workout
// @Produce      json
// @Failure  	 400
// @Failure  	 500
// @Success 	 200
// @Router       /workout/generate [post]
// @Security ApiKeyAuth
func GenerateWorkoutPlan(context *gin.Context) {
	var excerciseList []*models.Excercise
	result := database.Instance.Order("RANDOM()").Limit((rand.Int()%3 + 5)).Find(&excerciseList)
	if result.Error != nil {
		log.Printf("Error: %v", result.Error)
	}

	tokenString := context.GetHeader("Authorization")
	token := tokenService.ParseAccessToken(strings.TrimPrefix(tokenString, "Bearer "))

	workout := models.Workout{
		UserId:     token.UserId,
		Excercises: excerciseList,
	}

	record := database.Instance.Omit("Excercises.*").Save(&workout)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}

	context.JSON(200, WorkoutPlanDetails{Excercises: workout.Excercises, DateCreated: workout.CreatedAt})

}

func GetWorkoutPlanList(context *gin.Context) {
}

func FetchWorkout(context *gin.Context) {
	authorizationHeader := context.GetHeader("Authorization")
	token := tokenService.ParseAccessToken(authorizationHeader)

	workoutId, err := context.Params.Get("id")
	if !err {
		context.Abort()
		return
	}

	var workout models.Workout
	record := database.Instance.First(&workout, "id = ? and userid = ?", workoutId, token.UserId)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}

	context.JSON(200, workout)

}
