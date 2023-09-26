package models


type Todo struct {
	// gorm.Model
	Title  				string `json:"title"`
	Description 		string `json:"description"`
	Status 				int32 `json:"status"`
}

func (Todo) TableName() string {
    return "todo"
}