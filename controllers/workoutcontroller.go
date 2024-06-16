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
	Excercises  []*ExcerciseDetails
	DateCreated time.Time
	Id          string
	UserId      string
}

type ExcerciseDetails struct {
	Id           string
	Name         string
	Category     string
	Instructions string
}

type GeneratePlanRequest struct {
	Bodyparts []string `json:"bodyparts"`
}

// GenerateWorkout godoc
// @Summary      Generate workout plan
// @Tags         Workout
// @Param        details   body      GeneratePlanRequest  true  "User expectations"
// @Produce      json
// @Failure  	 400
// @Failure  	 500
// @Success 	 200
// @Router       /workout/generate [post]
// @Security ApiKeyAuth
func GenerateWorkoutPlan(context *gin.Context) {
	var requestDetails GeneratePlanRequest
	if err := context.ShouldBindJSON(&requestDetails); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var excerciseList []*models.Excercise
	result := database.Instance.Order("RANDOM()").Where("category IN ?", requestDetails.Bodyparts).Limit((rand.Int()%3 + 5)).Find(&excerciseList)
	if result.Error != nil {
		log.Printf("Error: %v", result.Error)
	}

	if len(excerciseList) == 0 {
		context.JSON(404, gin.H{"error": "No excercises found! Please change your criteria and try again"})
		context.Abort()
		return
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

	response := WorkoutPlanDetails{
		Excercises:  make([]*ExcerciseDetails, len(excerciseList)),
		DateCreated: workout.CreatedAt,
		UserId:      workout.UserId,
		Id:          workout.ID.String(),
	}

	for i, excercise := range excerciseList {
		response.Excercises[i] = &ExcerciseDetails{
			Id:           excercise.ID.String(),
			Category:     excercise.Category,
			Name:         excercise.Name,
			Instructions: excercise.Instructions,
		}
	}

	context.JSON(200, response)

}

// GetWorkouts godoc
// @Summary      Get all generated workout plans
// @Tags         Workout
// @Produce      json
// @Failure  	 400
// @Failure  	 500
// @Success 	 200
// @Router       /workout/all [get]
// @Security ApiKeyAuth
func GetWorkoutPlanList(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	token := tokenService.ParseAccessToken(strings.TrimPrefix(tokenString, "Bearer "))

	var workouts []*models.Workout
	record := database.Instance.Model(&models.Workout{}).Preload("Excercises").Where("user_id = ?", token.UserId).Find(&workouts)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}

	response := make([]*WorkoutPlanDetails, len(workouts))

	for i, plan := range workouts {
		response[i] = &WorkoutPlanDetails{
			DateCreated: plan.CreatedAt,
			UserId:      token.UserId,
			Id:          plan.ID.String(),
			Excercises:  make([]*ExcerciseDetails, len(plan.Excercises)),
		}

		for j, excercise := range plan.Excercises {
			response[i].Excercises[j] = &ExcerciseDetails{
				Id:           excercise.ID.String(),
				Category:     excercise.Category,
				Name:         excercise.Name,
				Instructions: excercise.Instructions,
			}
		}

	}

	context.JSON(200, response)
}

// Fetch workout godoc
// @Summary      Fetch single workout info
// @Tags         Workout
// @Param        uuid   path      string  true  "Workout ID"
// @Produce      json
// @Failure  	 400
// @Failure  	 500
// @Success 	 200
// @Router       /workout/{uuid} [get]
// @Security ApiKeyAuth
func FetchWorkout(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	token := tokenService.ParseAccessToken(strings.TrimPrefix(tokenString, "Bearer "))

	workoutId, err := context.Params.Get("id")
	if !err {
		context.Abort()
		return
	}

	var workout models.Workout
	record := database.Instance.Model(&models.Workout{}).Preload("Excercises").First(&workout, "id = ? and user_id = ?", workoutId, token.UserId)
	if record.Error != nil {
		context.JSON(400, gin.H{"error": record.Error})
		context.Abort()
		return
	}

	response := WorkoutPlanDetails{
		Excercises:  make([]*ExcerciseDetails, len(workout.Excercises)),
		DateCreated: workout.CreatedAt,
		UserId:      workout.UserId,
		Id:          workout.ID.String(),
	}

	for i, excercise := range workout.Excercises {
		response.Excercises[i] = &ExcerciseDetails{
			Id:           excercise.ID.String(),
			Category:     excercise.Category,
			Name:         excercise.Name,
			Instructions: excercise.Instructions,
		}
	}
	context.JSON(200, response)

}
