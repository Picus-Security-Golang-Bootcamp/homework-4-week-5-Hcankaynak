package author

import (
	"fmt"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (a *Author) ToString() string {
	return fmt.Sprintf("Id: %d \n"+
		"Name: %s \n", a.Id, a.Name)
}
