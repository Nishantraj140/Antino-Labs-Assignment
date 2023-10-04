package blog

import (
	"Antino-Labs-Assignment/pkg/sql"
)

type Article struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a *Article) Create() (err error) {
	err = sql.DB.Model(a).Create(a).Error
	return
}

func (a *Article) GetArticleByID() (err error) {
	err = sql.DB.Where("id = ?", a.ID).First(&a).Error
	return
}

func (a *Article) GetAllArticles() ([]Article, error) {
	var articles []Article
	err := sql.DB.Find(&articles).Error
	return articles, err
}

func (a *Article) Update() (err error) {
	err = sql.DB.Model(a).Updates(a).Error
	return err
}

func (a *Article) Delete() (err error) {
	err = sql.DB.Model(a).Delete(a).Error
	return err
}
