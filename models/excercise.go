package models

type Excercise struct {
	BaseModel
	Name      string      `json:"name"`
	BodyParts []*BodyPart `gorm:"many2many:excercise_bodyparts" json:"bodyparts"`
	Intensity int         `json:"intensity"`
}
