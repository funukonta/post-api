package services

import (
	"post-api/api/models"
	"post-api/api/repos"
)

type Post_Service interface {
	GetAll() ([]models.Post, error)
	GetById(*models.Post) (*models.PostRes, error)
	Create(*models.PostReq) (*models.PostRes, error)
	Update(int, *models.PostReq) error
	Delete(int) error
}

type post_Service struct {
	repo repos.Post_Repo
}

func New_PostService(repo repos.Post_Repo) Post_Service {
	return &post_Service{repo: repo}
}

func (s *post_Service) GetAll() ([]models.Post, error) {
	posts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (s *post_Service) GetById(post *models.Post) (*models.PostRes, error) {
	err := s.repo.GetById(post)
	if err != nil {
		return nil, err
	}

	// Formating response
	tagRes := []models.Tag{}
	for _, tag := range post.Tags {
		tagRes = append(tagRes, models.Tag{ID: tag.ID, Label: tag.Label})
	}
	postRes := &models.PostRes{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		Tags:    tagRes,
	}
	return postRes, nil
}

func (s *post_Service) Create(post *models.PostReq) (*models.PostRes, error) {
	var tags []models.Tag
	for _, label := range post.Tags {
		tags = append(tags, models.Tag{Label: label})
	}

	// formating db
	postModel := &models.Post{
		Title:   post.Title,
		Content: post.Content,
		Tags:    tags,
	}

	err := s.repo.Create(postModel)
	if err != nil {
		return nil, err
	}

	// formating response
	tagRes := []models.Tag{}
	for _, tag := range postModel.Tags {
		tagres := models.Tag{ID: tag.ID, Label: tag.Label}
		tagRes = append(tagRes, tagres)
	}
	postRes := &models.PostRes{ID: postModel.ID, Title: postModel.Title, Content: postModel.Content, Tags: tagRes}

	return postRes, nil

}

func (s *post_Service) Update(id int, postReq *models.PostReq) error {
	tags := []models.Tag{}
	for _, label := range postReq.Tags {
		tags = append(tags, models.Tag{Label: label})
	}

	err := s.repo.Update(&models.Post{ID: id, Title: postReq.Title, Content: postReq.Content, Tags: tags})
	if err != nil {
		return err
	}

	return nil
}

func (s *post_Service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
