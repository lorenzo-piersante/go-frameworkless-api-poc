package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/lorenzo-piersante/go-frameworkless-api-poc/storage"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type PostActionInput struct {
	Username string
	Password string
}

type PostActionOutput struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func (a *API) PostAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var input PostActionInput
	err := decoder.Decode(&input)

	if err != nil {
		respond(w, 400, []byte(`{"message":"failed decoding request body"}`))
		return
	}

	user, err := CreateUser(a, input)
	if err != nil || user == nil {
		respond(w, 500, []byte(`{"message":"internal server error"}`))
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(PostActionOutput{user.Id, user.Username})
	if err != nil {
		respond(w, 500, []byte(`{"message":"failed encoding response"}`))
		return
	}

	respond(w, 200, []byte(""))
}

func CreateUser(a *API, input PostActionInput) (*storage.User, error) {
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
