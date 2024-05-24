package models

type BodyPart struct {
	BaseModel
	Name       string       `json:"name"`
	Excercises []*Excercise `gorm:"many2many:bodypart_excercises" json:"excercises"`
}
