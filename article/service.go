package article

import (
	"blog.phumaster.dev-api/database"
	"blog.phumaster.dev-api/model"
)

// Service nil
type Service struct{}

// Find nil
func (a *Service) Find(id string) string {
	if id == "" {
		return "No ID provide"
	}
	return id
}

// All nil
func (a *Service) All(id string) string {
	if id == "" {
		return "No ID provide"
	}
	return id
}

// Create nil
func (a *Service) Create() model.Article {
	article := model.Article{
		Title:       "phumaster",
		Description: "yoooooooooo",
		Content:     "content",
		UserID:      1,
	}
	database.Connection.Create(&article)
	return article
}

// Update nil
func (a *Service) Update(id string) string {
	if id == "" {
		return "No ID provide"
	}
	return id
}

// Delete nil
func (a *Service) Delete(id string) string {
	if id == "" {
		return "No ID provide"
	}
	return id
}
