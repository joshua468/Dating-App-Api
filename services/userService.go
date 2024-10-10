package services

import (
	"errors"

	"github.com/joshua468/Dating-App-Api/internal/models"
	"github.com/joshua468/Dating-App-Api/internal/pkg/utils"
)

var users = make(map[string]models.User)
var matches = make([]models.Match, 0) // For storing matches

func CreateUser(user *models.User) error {
	if _, exists := users[user.Username]; exists {
		return errors.New("username already exists")
	}
	user.ID = uint(len(users) + 1)
	user.Password = utils.HashPassword(user.Password)
	users[user.Username] = *user
	return nil
}

func AuthenticateUser(username, password string) (models.User, error) {
	user, exists := users[username]
	if !exists || !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("authentication failed")
	}
	return user, nil
}

func CreateMatch(match *models.Match) error {
	if match.UserID == 0 || match.MatchedUserID == 0 {
		return errors.New("both user IDs are required")
	}
	matches = append(matches, *match) // Add match to the slice
	return nil
}

func GetMatches(userID uint) ([]models.Match, error) {
	var userMatches []models.Match
	for _, match := range matches {
		if match.UserID == userID {
			userMatches = append(userMatches, match)
		}
	}
	if len(userMatches) == 0 {
		return nil, errors.New("no matches found")
	}
	return userMatches, nil
}
