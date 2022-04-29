package userRegist

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Name         string
	Email        string
	Phone_number string
	Created_at   time.Time
	Updated_at   time.Time
	Is_deleted   int
}

func InsertUser(db *gorm.DB, values map[string]string) {
	user := User{
		Name:         values["user_name"],
		Email:        values["email"],
		Phone_number: values["phone_number"],
		Created_at:   time.Now().Add(time.Hour * 9),
		Updated_at:   time.Now().Add(time.Hour * 9),
		Is_deleted:   0}

	db.Create(&user)
	return
}
