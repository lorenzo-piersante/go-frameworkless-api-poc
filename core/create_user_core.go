package core

import (
	"github.com/google/uuid"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/api"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(a *api.API, input api.PostActionInput) (*storage.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}

	user := storage.User{
		Id:       uuid.New().String(),
		Username: input.Username,
		Password: string(hashedPassword),
	}

	err = a.Storage.StoreUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
