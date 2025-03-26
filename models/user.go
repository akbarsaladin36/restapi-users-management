package models

import "time"

type User struct {
	UserId                  int       `json:"user_id" gorm:"primaryKey"`
	UserUuid                string    `json:"user_uuid" gorm:"type:varchar(200)"`
	UserUsername            string    `json:"user_username" gorm:"type:varchar(100)"`
	UserEmail               string    `json:"user_email" gorm:"type:varchar(150)"`
	UserPassword            string    `json:"user_password" gorm:"type:varchar(200)"`
	UserFirstName           string    `json:"user_first_name" gorm:"type:varchar(150)"`
	UserLastName            string    `json:"user_last_name" gorm:"type:varchar(150)"`
	UserAddress             string    `json:"user_address" gorm:"type:text"`
	UserPhoneNumber         string    `json:"user_phone_number" gorm:"type:varchar(30)"`
	UserStatusCd            string    `json:"user_status_cd" gorm:"type:varchar(30)"`
	UserRole                string    `json:"user_role" gorm:"type:varchar(30)"`
	UserCreatedDate         time.Time `json:"user_created_date"`
	UserCreatedUserUuid     string    `json:"user_created_user_uuid" gorm:"type:varchar(200)"`
	UserCreatedUserUsername string    `json:"user_created_user_username" gorm:"type:varchar(100)"`
	UserUpdatedDate         time.Time `json:"user_updated_date"`
	UserUpdatedUserUuid     string    `json:"user_updated_user_uuid" gorm:"type:varchar(200)"`
	UserUpdatedUserUsername string    `json:"user_updated_user_username" gorm:"type:varchar(100)"`
}
