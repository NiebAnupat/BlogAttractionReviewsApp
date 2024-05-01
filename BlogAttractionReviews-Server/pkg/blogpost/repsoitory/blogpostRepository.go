package repsoitory

import "github.com/NiebAnupat/BlogAttractionReviewsApp/Server/entities"

type BlogPostRepository interface {
	Create(blogPostEntity *entities.BlogPost) (string, error)
	FindByID(id string) (*entities.BlogPost, error)
	FindAll() ([]entities.BlogPost, error)
	Delete(id string) error
}
