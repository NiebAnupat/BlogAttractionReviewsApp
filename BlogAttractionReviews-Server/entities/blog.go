package entities

import (
	"time"

	"github.com/NiebAnupat/BlogAttractionReviewsApp/Server/pkg/blogpost/model"
)

type (
	BlogPost struct {
		ID            string         `gorm:"primaryKey;type:varchar(64);"`
		Title         string         `gorm:"type:varchar(64);not null;"`
		Description   string         `gorm:"type:text;not null;"`
		Thumbnail     string         `gorm:"type:varchar(64);not null;"`
		CreateAt      time.Time      `gorm:"not null;autoCreateTime;"`
		UpdateAt      time.Time      `gorm:"not null;autoUpdateTime;"`
		Contents      []*BlogContent `gorm:"foreignKey:BlogID"`
		AuthorID      string         `gorm:"type:varchar(64);not null;"`
		UserLikes     []UserLike     `gorm:"foreignKey:BlogID"`
		UserFavorites []UserFavorite `gorm:"foreignKey:BlogID"`
	}

	BlogContent struct {
		ID     string `gorm:"primaryKey;type:varchar(64);"`
		BlogID string `gorm:"type:varchar(64);not null;"`
		Order  int    `gorm:"not null;"`
		Type   int    `gorm:"not null; default:0; description:'0=text, 1=image';"`
		Value  string `gorm:"type:text;null;"`
	}
)

func (b *BlogPost) ToBlogPostModel() *model.BlogPost {
	return &model.BlogPost{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Thumbnail:   b.Thumbnail,
		CreateAt:    b.CreateAt,
		AuthorID:    b.AuthorID,
		Contents:    b.ToBlogContentModel(),
		Likes:       len(b.UserLikes),
		Favorites:   len(b.UserFavorites),
	}
}

func (b *BlogPost) ToBlogContentModel() []*model.BlogContent {
	var blogContents []*model.BlogContent
	for _, content := range b.Contents {
		blogContents = append(blogContents, &model.BlogContent{
			ID:    content.ID,
			Order: content.Order,
			Type:  content.Type,
			Value: content.Value,
		})
	}
	return blogContents
}

func (c *BlogContent) ToBlogContentModel() *model.BlogContent {
	return &model.BlogContent{
		ID:    c.ID,
		Order: c.Order,
		Type:  c.Type,
		Value: c.Value,
	}
}
