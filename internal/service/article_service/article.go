package article

import (
	"encoding/json"

	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/gredis"
	"Go-blog-server/pkg/logging"
	"Go-blog-server/internal/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

// func (a *Article) Add() error {
// 	article := map[string]interface{}{
// 		"tag_id":          a.TagID,
// 		"title":           a.Title,
// 		"desc":            a.Desc,
// 		"content":         a.Content,
// 		"created_by":      a.CreatedBy,
// 		"cover_image_url": a.CoverImageUrl,
// 		"state":           a.State,
// 	}

// 	if err := models.AddArticle(article); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (a *Article) Edit() error {
// 	return models.EditArticle(a.ID, map[string]interface{}{
// 		"tag_id":          a.TagID,
// 		"title":           a.Title,
// 		"desc":            a.Desc,
// 		"content":         a.Content,
// 		"cover_image_url": a.CoverImageUrl,
// 		"state":           a.State,
// 		"modified_by":     a.ModifiedBy,
// 	})
// }

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}