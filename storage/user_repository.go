package storage

import (
	"encoding/json"
	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
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
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
