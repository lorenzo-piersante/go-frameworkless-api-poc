package core

import (
	"github.com/google/uuid"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(s *storage.Storage, username string, password string) (*storage.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user := storage.User{
		Id:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
	}

	err = s.StoreUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
