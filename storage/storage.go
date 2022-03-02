package storage

import (
	"encoding/json"
	"fmt"
)

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

func (s *Storage) GetUserByID(id string) (p *User, err error) {
	b, err := s.redis.Get(dbKey + id).Bytes()
	if err != nil {
		err = fmt.Errorf("failed to get page %s: %v", id, err)
		return
	}

	err = json.Unmarshal(b, &p)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal %s: %v", id, err)
		return
	}

	return
}
