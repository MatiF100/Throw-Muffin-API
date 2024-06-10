package models

type Excercise struct {
	BaseModel
	Name         string `json:"name"`
	Category     string `json:"category"`
	Instructions string `json:"instructions"`
}
