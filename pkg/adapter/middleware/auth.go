package middleware

import (
	"context"
	"log"
	"net/http"

	"22dojo-online/pkg/adapter/dcontext"
	"22dojo-online/pkg/errors"
	"22dojo-online/pkg/usecase"
)

type Auth struct {
	userInteractor usecase.UserInteractorInterface
}

func NewAuth(userInteractor usecase.UserInteractorInterface) *Auth {
	return &Auth{
		userInteractor: userInteractor,
	}
}

func (auth *Auth) Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// リクエストヘッダからx-token(認証トークン)を取得
		token := request.Header.Get("x-token")
		if token == "" {
			log.Println("x-token is empty")
			return
		}

		user, err := auth.userInteractor.SelectUserByAuthToken(token)
		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "Internal Server Error")
			return
		}

		if err != nil {
			log.Println(err)
			errors.InternalServerError(writer, "Invalid token")
			return
		}
		if user == nil {
			log.Printf("user not found. token=%s", token)
			errors.BadRequest(writer, "Invalid token")
			return
		}

		// ユーザIDをContextへ保存して以降の処理に利用する
		ctx = dcontext.SetUserID(ctx, user.ID)

		// 次の処理
		nextFunc(writer, request.WithContext(ctx))
	}
}
