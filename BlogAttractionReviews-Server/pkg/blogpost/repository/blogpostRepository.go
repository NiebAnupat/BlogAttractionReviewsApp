package repository

import "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"

type BlogPostRepository interface {
	Create(blogPostEntity *entities.BlogPost) (*entities.BlogPost, error)
	CreateContent(blogContentEntity *entities.BlogContent) (*entities.BlogContent, error)
	FindByID(id string) (*entities.BlogPost, error)
	FindAll() ([]entities.BlogPost, error)
	FindAllByAuthorID(authorID string) ([]entities.BlogPost, error)
	Delete(id string) error
}
