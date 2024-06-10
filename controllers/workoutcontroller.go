package controllers

import (
	"fmt"
	"math/rand"

	"github.com/MatiF100/Throw-Muffin-API/database"
	"github.com/MatiF100/Throw-Muffin-API/models"
	"github.com/gin-gonic/gin"
)

func GenerateWorkoutPlan(context *gin.Context) {
	var excerciseList []models.Excercise
	result := database.Instance.Order("rand()").Limit((rand.Int()%3 + 1) * 10).Find(&excerciseList)
	if result.Error != nil {

	}
	for _, ex := range excerciseList {
		fmt.Println(ex)
	}
}

func GetWorkoutPlanList(context *gin.Context) {
}

func FetchWorkout(context *gin.Context) {

}
