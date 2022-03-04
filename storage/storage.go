package storage

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
)

var ctx = context.Background()

type Storage struct {
	redis *redis.Client
}

func NewStorage(redis *redis.Client) *Storage {
	return &Storage{
		redis,
	}
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"title"`
	Password string `json:"password"`
}

func (s *Storage) GetUserById(id string) (user *User, err error) {
	u, err := s.redis.Get(ctx, id).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(u, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Storage) CreateUser(username string, password string) (*User, error) {
	user := User{
		uuid.New().String(),
		username,
		password,
	}

	err := s.redis.Set(ctx, user.Id, user, 0).Err()
	log.Fatal(err)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
