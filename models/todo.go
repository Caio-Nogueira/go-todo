package models

import (
	"time"
)

type Todo struct {
	// gorm.Model
	Title  				string `json:"title"`
	Description 		string `json:"description"`
	Status 				int32 `json:"status"`
	UserId 				string `json:"user_id" gorm:"foreignKey:user_id;type:uuid"`
	CreatedAt 			time.Time `json:"created_at"`
	UpdatedAt 			time.Time `json:"updated_at"`
}

func (Todo) TableName() string {
    return "todo"
}