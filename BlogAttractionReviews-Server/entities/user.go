package entities

import "time"

type (
	User struct {
		ID       string `gorm:"primaryKey;type:varchar(64);"`
		Username string `gorm:"type:varchar(64);not null;unique"`
		Email    string `gorm:"type:varchar(64);not null;unique"`
		// Password string `gorm:"type:varchar(64);not null"`
		Avatar        string         `gorm:"type:varchar(64);not null"`
		CreateAt      time.Time      `gorm:"not null;autoCreateTime;"`
		UpdateAt      time.Time      `gorm:"not null;autoUpdateTime;"`
		UserLikes     []UserLike     `gorm:"foreignKey:UserID"`
		UserFavorites []UserFavorite `gorm:"foreignKey:UserID"`
	}

	UserLike struct {
		ID     string `gorm:"primaryKey;type:varchar(64);"`
		UserID string `gorm:"type:varchar(64);not null;"`
		BlogID string `gorm:"type:varchar(64);not null;"`
	}

	UserFavorite struct {
		ID     string `gorm:"primaryKey;type:varchar(64);"`
		UserID string `gorm:"type:varchar(64);not null;"`
		BlogID string `gorm:"type:varchar(64);not null;"`
	}
)
