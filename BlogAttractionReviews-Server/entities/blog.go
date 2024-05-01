package entities

type (
	BlogPost struct {
		ID            string         `gorm:"primaryKey;type:varchar(64);"`
		Title         string         `gorm:"type:varchar(64);not null;unique"`
		CreateAt      string         `gorm:"not null;autoCreateTime;"`
		UpdateAt      string         `gorm:"not null;autoUpdateTime;"`
		Contents      []BlogContent  `gorm:"foreignKey:BlogID"`
		AuthorID      string         `gorm:"type:varchar(64);not null;"`
		UserLikes     []UserLike     `gorm:"foreignKey:BlogID"`
		UserFavorites []UserFavorite `gorm:"foreignKey:BlogID"`
	}

	BlogContent struct {
		ID       string `gorm:"primaryKey;type:varchar(64);"`
		BlogID   string `gorm:"type:varchar(64);not null;"`
		Order    int    `gorm:"not null;"`
		Type     int    `gorm:"not null; default:0; description:'0=text, 1=image';"`
		Text     string `gorm:"type:text;null;"`
		ImageURL string `gorm:"type:varchar(64);null;"`
	}
)
