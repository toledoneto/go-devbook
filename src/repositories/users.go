package repositories

import (
	"api/src/models"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Users struct {
	db *gorm.DB
}

type Followers struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {

	result := repository.db.Create(&user)
	if result.Error != nil {
		return 0, nil
	}

	userID := user.ID

	return uint64(userID), nil
}

func (repository Users) Search(searchPattern string) ([]models.User, error) {
	searchPattern = fmt.Sprintf("%%%s%%", searchPattern)

	results, err := repository.db.Model(&models.User{}).
		Select("id", "name", "username", "email").
		Where(
			repository.db.Where("name ILIKE ?", searchPattern)).
		Or(
			repository.db.Where("username ILIKE ?", searchPattern)).
		Rows()
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var users []models.User
	for results.Next() {
		var user models.User

		repository.db.ScanRows(results, &user)

		users = append(users, user)

	}

	return users, nil
}

func (repository Users) SearchUser(userID uint64) (models.User, error) {
	var user models.User

	result := repository.db.Model(&models.User{}).
		Select("id", "name", "username", "email").
		First(&user, "id = ?", userID)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (repository Users) UpdateUser(userID uint64, user models.User) error {

	user = models.User{Name: user.Name,
		Username: user.Username,
		Email:    user.Email}

	result := repository.db.Model(&user).
		Where("id = ?", userID).
		Updates(models.User{Name: user.Name,
			Username: user.Username,
			Email:    user.Email})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Users) DeleteUser(userID uint64) error {

	if result := repository.db.Model(&models.User{}).
		Where("id = ?", userID).
		Delete(&models.User{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Users) SearchUserByEmail(email string) (models.User, error) {
	var user models.User

	result := repository.db.Model(&models.User{}).
		Select("id", "password").
		First(&user, "email = ?", email)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil

}

func (repository Users) Follow(userID, followerID uint64) error {
	follow := models.Followers{User_ID: userID, Follower_ID: followerID}
	result := repository.db.
		Select("user_id", "follower_id").
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&follow)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Users) Unfollow(userID, followerID uint64) error {
	if result := repository.db.Model(&models.Followers{}).
		Where("user_id = ?", userID).
		Where("follower_id = ?", followerID).
		Delete(&models.Followers{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	results, err := repository.db.Model(&models.User{}).
		Select("users.id, users.name, users.email, users.username, users.created_at").
		Joins("inner join followers on users.id = followers.follower_id").
		Where("followers.user_id = ?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var users []models.User
	for results.Next() {
		var user models.User
		repository.db.ScanRows(results, &user)
		users = append(users, user)

	}
	return users, nil
}

func (repository Users) SearchFollowing(userID uint64) ([]models.User, error) {
	results, err := repository.db.Model(&models.User{}).
		Select("users.id, users.name, users.email, users.username, users.created_at").
		Joins("inner join followers on users.id = followers.user_id").
		Where("followers.follower_id = ?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var users []models.User
	for results.Next() {
		var user models.User
		repository.db.ScanRows(results, &user)
		users = append(users, user)

	}
	return users, nil
}

func (repository Users) FetchPassword(userID uint64) (string, error) {
	var user models.User

	result := repository.db.Model(&models.User{}).
		Select("password").
		First(&user, "id = ?", userID)
	if result.Error != nil {
		return "", result.Error
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userID uint64, password string) error {

	result := repository.db.Model(&models.User{}).
		Where("id = ?", userID).
		Updates(models.User{
			Password: password,
		})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
