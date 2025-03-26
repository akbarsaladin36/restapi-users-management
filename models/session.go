package models

import "time"

type Session struct {
	SessionId                  string    `json:"session_id" gorm:"type:varchar(150)"`
	SessionToken               string    `json:"session_token" gorm:"type:text"`
	SessionUserUuid            string    `json:"session_user_uuid" gorm:"type:varchar(200)"`
	SessionUserUsername        string    `json:"session_user_username" gorm:"type:varchar(100)"`
	SessionUserRole            string    `json:"session_user_role" gorm:"type:varchar(30)"`
	SessionStartAt             time.Time `json:"session_start_at"`
	SessionExpiredAt           time.Time `json:"session_expired_at"`
	SessionStatusCd            string    `json:"session_status_cd" gorm:"type:varchar(30)"`
	SessionCreatedDate         time.Time `json:"session_created_date"`
	SessionCreatedUserUuid     string    `json:"session_created_user_uuid" gorm:"type:varchar(200)"`
	SessionCreatedUserUsername string    `json:"session_created_user_username" gorm:"type:varchar(100)"`
	SessionUpdatedDate         time.Time `json:"session_updated_date"`
	SessionUpdatedUserUuid     string    `json:"session_updated_user_uuid" gorm:"type:varchar(200)"`
	SessionUpdatedUserUsername string    `json:"session_updated_user_username" gorm:"type:varchar(100)"`
}
