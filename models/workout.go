package models

type Workout struct {
	BaseModel
	Excercises []Excercise `json:"excercises" gorm:"many2many:workout_excercises"`
	UserId     string      `json:"userId"`
}
