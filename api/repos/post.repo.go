package repos

import (
	"fmt"
	"post-api/api/models"

	"gorm.io/gorm"
)

type Post_Repo interface {
	GetAll() ([]models.Post, error)
	GetById(*models.Post) error
	Create(*models.Post) error
	Update(*models.Post) error
	Delete(id int) error
}

type post_Repo struct {
	db *gorm.DB
}

func New_PostRepo(db *gorm.DB) Post_Repo {
	return &post_Repo{db: db}
}

func (r *post_Repo) GetAll() ([]models.Post, error) {
	postsDb := []models.Post{}
	err := r.db.Preload("Tags").Find(&postsDb).Error
	if err != nil {
		return nil, err
	}

	return postsDb, nil
}

func (r *post_Repo) GetById(post *models.Post) error {
	err := r.db.Preload("Tags").First(post, "id = ?", post.ID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("post tidak ditemukan")
		}

		return err
	}
	return nil
}

func (r *post_Repo) Create(post *models.Post) error {
	// tags := []models.Tag{}
	for i, tag := range post.Tags {
		// tagRes := models.Tag{}
		err := r.db.FirstOrCreate(&tag, models.Tag{Label: tag.Label}).Error
		if err != nil {
			return err
		}
		post.Tags[i] = tag
		// tags = append(tags, tagRes)
	}

	err := r.db.Create(post).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *post_Repo) Update(postReq *models.Post) error {
	var tags []models.Tag
	for _, tag := range postReq.Tags {
		tagDb := models.Tag{Label: tag.Label}
		err := r.db.Where("label = ?", tagDb.Label).First(&tagDb).Error
		if err == gorm.ErrRecordNotFound {
			err := r.db.Create(&tagDb).Error
			if err != nil {
				return err
			}
		}
		tags = append(tags, tagDb)
	}

	err := r.db.Model(&postReq).Association("Tags").Replace(tags)
	if err != nil {
		return err
	}

	err = r.db.Save(&postReq).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *post_Repo) Delete(id int) error {
	var post models.Post

	err := r.db.Preload("Tags").First(&post, id).Error
	if err != nil {
		return err
	}

	err = r.db.Model(&post).Association("Tags").Clear()
	if err != nil {
		return err
	}

	err = r.db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}
