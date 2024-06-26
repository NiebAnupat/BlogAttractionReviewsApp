package entities

import (
	"time"

	userModel "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/user/model"
)

type (
	User struct {
		ID            string         `gorm:"primaryKey;type:varchar(64);"`
		Username      string         `gorm:"type:varchar(64);not null;unique"`
		Password      string         `gorm:"type:varchar(64);not null"`
		Avatar        string         `gorm:"type:varchar(64);not null"`
		CreateAt      time.Time      `gorm:"not null;autoCreateTime;"`
		UpdateAt      time.Time      `gorm:"not null;autoUpdateTime;"`
		UserLikes     []UserLike     `gorm:"foreignKey:UserID"`
		UserFavorites []UserFavorite `gorm:"foreignKey:UserID"`
		BlogPosts     []BlogPost     `gorm:"foreignKey:AuthorID"`
	}

	UserLike struct {
		ID     string `gorm:"primaryKey;type:varchar(64);"`
		UserID string `gorm:"type:varchar(64);not null;"`
		BlogID string `gorm:"type:varchar(64);not null;"`
		// User   User   `gorm:"foreignKey:UserID"`
	}

	UserFavorite struct {
		ID     string `gorm:"primaryKey;type:varchar(64);"`
		UserID string `gorm:"type:varchar(64);not null;"`
		BlogID string `gorm:"type:varchar(64);not null;"`
	}
)

func (u *User) ToUserModel() *userModel.User {
	return &userModel.User{
		ID:       u.ID,
		Username: u.Username,
		Avatar:   u.Avatar,
	}
}
