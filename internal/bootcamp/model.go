package bootcamp

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}
