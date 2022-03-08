package storage

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Storage) GetUserById(id string) (user *User, err error) {
	rows, err := s.db.Query("select users.id, users.username, users.password FROM users WHERE users.id = ?")
	if err != nil {
		return nil, err
	}

	var username string
	var password string
	err = rows.Scan(&username, &password)
	if err != nil {
		return nil, err
	}

	return &User{id, username, password}, nil
}

func (s *Storage) StoreUser(user User) error {
	statement, err := s.db.Prepare("INSERT INTO users (id, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(user.Id, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
