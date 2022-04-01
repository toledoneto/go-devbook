package repositories

import (
	"api/src/models"

	"gorm.io/gorm"
)

type Posts struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {

	result := repository.db.
		Select("title", "content", "author_id").
		Create(&post)
	if result.Error != nil {
		return 0, nil
	}

	postID := post.ID

	return uint64(postID), nil
}

func (repository Posts) SearchPost(postID uint64) (models.Post, error) {
	var post models.Post
	result := repository.db.Model(&models.Post{}).
		Select("posts.*, users.username").
		Joins("inner join users on users.id = posts.author_id").
		Where("posts.id = ?", postID).
		Scan(&post)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return post, nil
}

func (repository Posts) SearchAllPosts(userID uint64) ([]models.Post, error) {
	results, err := repository.db.Model(&models.Post{}).
		Distinct("posts.*, users.username").
		Joins("inner join users on users.id = posts.author_id").
		Joins("inner join followers on followers.user_id = posts.author_id").
		Where(
			repository.db.Where("users.id = ?", userID)).
		Or(
			repository.db.Where("followers.follower_id = ?", userID)).
		Order("1 desc").
		Rows()
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var posts []models.Post

	for results.Next() {
		var post models.Post

		repository.db.ScanRows(results, &post)

		posts = append(posts, post)

	}

	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	result := repository.db.Model(&models.Post{}).
		Where("id = ?", postID).
		Updates(models.Post{Title: post.Title,
			Content: post.Content})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	if result := repository.db.Model(&models.Post{}).
		Where("id = ?", postID).
		Delete(&models.Post{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Posts) SearchPostsByUser(userID uint64) ([]models.Post, error) {
	results, err := repository.db.Model(&models.Post{}).
		Select("posts.*, users.username").
		Joins("inner join users on users.id = posts.author_id").
		Where("posts.author_id = ?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var posts []models.Post

	for results.Next() {
		var post models.Post

		repository.db.ScanRows(results, &post)

		posts = append(posts, post)

	}

	return posts, nil
}

func (repository Posts) Like(postID uint64) error {
	result := repository.db.Model(&models.Post{}).
		Where("id = ?", postID).
		Update("likes", gorm.Expr("likes + ?", 1))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Posts) Unlike(postID uint64) error {
	result := repository.db.Model(&models.Post{}).
		Where("id = ?", postID).
		Where("likes > 0").
		Update("likes", gorm.Expr("likes - ?", 1))
	if result.Error != nil {
		return result.Error
	}

	return nil
}
