package domain

import (
	"time"
)

type User struct {
	Id           int       `json:"ID" gorm:"primaryKey"`
	UserId       string    `json:"USER_ID" gorm:"unique;not null"`
	Password     string    `json:"-"	gorm:"not null"` // JSON 응답에서 제외
	UserName     string    `json:"USER_NAME" gorm:"not null"`
	UserStatus   bool      `json:"USER_STATUS" gorm:"not null"`
	Email        string    `json:"EMAIL" gorm:"unique;not null"`
	PhoneNumber  string    `json:"PHONE_NUMBER" gorm:"not null"`
	JobCode      int       `json:"JOB_CODE"`
	ProfileImage string    `json:"PROFILE_IMAGE"`
	GithubLink   string    `json:"GITHUB_LINK"`
	BlogLink     string    `json:"BLOG_LINK"`
	UserText     string    `json:"USER_TEXT"`
	Company      string    `json:"COMPANY"`
	Skill        string    `json:"SKILL"`
	CreateDate   time.Time `json:"CREATE_DATE"`
	UpdateDate   time.Time `json:"UPDATE_DATE"`
}
