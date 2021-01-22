package foodratingmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "FoodRating"

// FoodRating model
type FoodRating struct {
	common.SQLModel
	UserID  int     `json:"user_id" gorm:"not null"`
	FoodID  int     `json:"food_id" gorm:"not null"`
	Point   float32 `json:"point" gorm:"default:0;not null"`
	Comment string  `json:"comment"`
}

// TableName table name
func (FoodRating) TableName() string {
	return "food_ratings"
}