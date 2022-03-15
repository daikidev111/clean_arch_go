package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"22dojo-online/pkg/adapter/dcontext"
	"22dojo-online/pkg/domain/entity"
	"22dojo-online/pkg/errors"
	"22dojo-online/pkg/usecase"

	"github.com/google/uuid"
)

type UserController struct {
	Interactor usecase.UserInteractorInterface
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

type UserGetResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	HighScore int32  `json:"highScore"`
	Coin      int32  `json:"coin"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

func NewUserController(ui usecase.UserInteractorInterface) *UserController {
	return &UserController{
		Interactor: ui,
	}
}

func (controller *UserController) GetUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if userID == "" {
			log.Println("userID is empty")
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}

		user, err := controller.Interactor.SelectUserByPrimaryKey(userID)
		if err != nil {
			log.Printf("[ERROR] GetUser() err = %s", err.Error())
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}
		if user == nil {
			log.Println("user not found")
			errors.BadRequest(writer, fmt.Sprintf("user not found. userID=%s", userID))
			return
		}

		// userHandler := handler.New()
		// userHandler.GetUserHandler(writer, user)
		body := &UserGetResponse{
			ID:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		}

		data, err := json.Marshal(body)
		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "marshal error")
			return
		}
		if _, err := writer.Write(data); err != nil {
			log.Println(err)
		}
	}
}

func (controller *UserController) InsertUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody UserCreateRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			errors.BadRequest(writer, "Bad Request")
			return
		}

		// TODO: request時にvalidator を使ってvalidation
		// UUIDでユーザIDを生成する
		userID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}

		// UUIDで認証トークンを生成する
		authToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}
		// データベースにユーザデータを登録する
		user := &entity.User{
			ID:        userID.String(),
			AuthToken: authToken.String(),
			Name:      requestBody.Name,
			HighScore: 0,
			Coin:      0,
		}
		if err := controller.Interactor.InsertUser(user); err != nil {
			log.Printf("[ERROR] InsertUser() err = %s", err.Error())
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}

		// userHandler := handler.New()
		// userHandler.CreateUserHandler(writer, authToken.String())
		body := &UserCreateResponse{
			Token: user.AuthToken,
		}
		data, err := json.Marshal(body)
		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "marshal error")
			return
		}
		if _, err := writer.Write(data); err != nil {
			log.Println(err)
		}
	}
}

func (controller *UserController) UpdateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody UserUpdateRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			errors.BadRequest(writer, "Bad Request")
			return
		}

		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if userID == "" {
			log.Println("userID is empty")
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}

		user, err := controller.Interactor.SelectUserByPrimaryKey(userID)
		if err != nil {
			log.Printf("[ERROR] UpdateUser() err = %s", err.Error())
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}
		if user == nil {
			log.Println("user not found")
			errors.BadRequest(writer, fmt.Sprintf("user not found. userID=%s", userID))
			return
		}
		user.Name = requestBody.Name

		err = controller.Interactor.UpdateUserByPrimaryKey(user)
		if err != nil {
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
